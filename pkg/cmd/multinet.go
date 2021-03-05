/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd/api"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/printers"

	netutils "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/utils"
)

var (
	podnetExample = `
	# view the pod network status in text (default)
        %[1]s 
	# view the pod network status in namespace, namespace=foobar
        %[1]s -n foobar
	# view the pod network status in json
        %[1]s -o json
	# view the pod network status in json, namespace=foobar
        %[1]s -o json -n foobar
`
)

// PodnetOptions provides ...
type PodnetOptions struct {
	configFlags *genericclioptions.ConfigFlags

	rawConfig api.Config
	args      []string

	allNamespaces bool
	namespace     string
	outputTarget  string
	outputFormat  string
	genericclioptions.IOStreams
}

// NewPodnetOptions provides an instance of NamespaceOptions with default values
func NewPodnetOptions(streams genericclioptions.IOStreams) *PodnetOptions {
	return &PodnetOptions{
		configFlags: genericclioptions.NewConfigFlags(true),

		IOStreams: streams,
	}
}

// NewCmdPodnet provides a cobra command wrapping NamespaceOptions
func NewCmdPodnet(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewPodnetOptions(streams)

	cmd := &cobra.Command{
		Use:          "podnet [flags]",
		Short:        "View network-status annotation of Pod",
		Example:      fmt.Sprintf(podnetExample, "kubectl"),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Complete(c, args); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}

			return o.Run()
		},
	}

	// add additional option if needed
	//cmd.Flags().BoolVar(&o.listNamespaces, "list", o.listNamespaces, "if true, print the list of all namespaces in the current KUBECONFIG")
	cmd.Flags().BoolVarP(&o.allNamespaces, "all-namespaces", "A", o.allNamespaces, "Show resources in all namespaces")
	cmd.Flags().StringVarP(&o.outputFormat, "output", "o", o.outputFormat, "Output format. One of: json|text|wide")
	//cmd.Flags().StringVarP(&o.outputTarget, "target", "t", o.outputTarget, "Output target. One of: default|memif|pci|vdpa|vhostuser")
	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}

// Complete sets all information required for updating the current context
func (o *PodnetOptions) Complete(cmd *cobra.Command, args []string) error {
	o.args = args

	var err error
	o.rawConfig, err = o.configFlags.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return err
	}

	o.namespace, err = cmd.Flags().GetString("namespace")
	if err != nil {
		return err
	}

	if o.namespace == "" {
		o.namespace = "default"
	}

	if o.allNamespaces {
		o.namespace = ""
	}

	return nil
}

// Validate ensures that all required arguments and flag values are provided
func (o *PodnetOptions) Validate() error {
	/*
		if len(o.rawConfig.CurrentContext) == 0 {
			return errNoContext
		}
		if len(o.args) > 1 {
			return fmt.Errorf("either one or no arguments are allowed")
		}
	*/

	if o.outputFormat == "" {
		o.outputFormat = "text"
	} else {
		o.outputFormat = strings.ToLower(o.outputFormat)

		switch o.outputFormat {
		case "json", "text", "wide": // valid format
		default: // illegal format
			return fmt.Errorf("unknown output format %s", o.outputFormat)
		}
	}

	return nil
}

// Run lists all available namespaces on a user's KUBECONFIG or updates the
// current context based on a provided namespace.
func (o *PodnetOptions) Run() error {
	switch o.outputTarget {
	/*
			//TBD
		case "pci":
			return o.ShowPCIOutput()
		case "vdpa":
			return o.ShowVdpaOutput()
		case "vhostuser":
			return o.ShowVhostUserOutput()
		case "memif":
			return o.ShowMemifOutput()
	*/
	default:
		return o.ShowDefaultOutput()
	}
}

// PodNetDefaultOutput is intermediate representation of output.
type PodNetDefaultOutput struct {
	Namespace string   `json:"namespace"`
	Pod       string   `json:"pod"`
	Net       string   `json:"net"`
	Interface string   `json:"interface"`
	Address   []string `json:"address"`
	Mac       string   `json:"mac"`
}

// ConvertRow convets PodNetDefaultOutput to TableRow's interface structure
func (p *PodNetDefaultOutput) ConvertRow() []interface{} {
	addrs := []interface{}{}
	for _, addr := range p.Address {
		addrs = append(addrs, addr)
	}
	return []interface{}{
		p.Pod,
		p.Interface,
		addrs,
		p.Mac,
	}
}

// ConvertWideRow convets PodNetDefaultOutput to TableRow's interface structure
func (p *PodNetDefaultOutput) ConvertWideRow() []interface{} {
	addrs := []interface{}{}
	for _, addr := range p.Address {
		addrs = append(addrs, addr)
	}
	return []interface{}{
		p.Namespace,
		p.Pod,
		p.Net,
		p.Interface,
		addrs,
		p.Mac,
	}
}

// ShowDefaultOutputJSON outputs PodNetDefaultOutput in json
func (o *PodnetOptions) ShowDefaultOutputJSON(output []PodNetDefaultOutput) error {
	outputJSON, err := json.Marshal(output)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "%s\n", outputJSON)
	return nil
}

// ShowDefaultOutputText outputs TableRow, converted from PodNetDefaultOutput, in text
func (o *PodnetOptions) ShowDefaultOutputText(rows []metav1.TableRow) error {
	// Table output
	columns := []metav1.TableColumnDefinition{
		{Name: "Pod", Type: "string"},
		{Name: "IF", Type: "string"},
		{Name: "Address", Type: "array"},
		{Name: "Mac", Type: "string"},
	}

	table := &metav1.Table{
		ColumnDefinitions: columns,
		Rows:              rows,
	}
	printer := printers.NewTablePrinter(printers.PrintOptions{})
	printer.PrintObj(table, os.Stdout)
	return nil
}

// ShowWideOutputText outputs TableRow, converted from PodNetWideOutput, in wide text
func (o *PodnetOptions) ShowWideOutputText(rows []metav1.TableRow) error {
	// Table output
	columns := []metav1.TableColumnDefinition{
		{Name: "Namespace", Type: "string"},
		{Name: "Pod", Type: "string"},
		{Name: "Net", Type: "string"},
		{Name: "IF", Type: "string"},
		{Name: "Address", Type: "array"},
		{Name: "Mac", Type: "string"},
	}

	table := &metav1.Table{
		ColumnDefinitions: columns,
		Rows:              rows,
	}
	printer := printers.NewTablePrinter(printers.PrintOptions{})
	printer.PrintObj(table, os.Stdout)
	return nil
}

// ShowDefaultOutput shows default (non DeviceInfo result)
func (o *PodnetOptions) ShowDefaultOutput() error {
	config, err := o.configFlags.ToRESTConfig()
	if err != nil {
		return err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	podList, err := client.CoreV1().Pods(o.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	outputs := []PodNetDefaultOutput{}

	rows := []metav1.TableRow{}
	for _, pod := range podList.Items {
		statuses, _ := netutils.GetNetworkStatus(&pod)
		for _, s := range statuses {
			// skip if device info case
			if s.DeviceInfo != nil {
				continue
			}
			p := PodNetDefaultOutput{}
			if pod.Namespace == "" {
				p.Namespace = "default"
			} else {
				p.Namespace = pod.Namespace
			}
			p.Pod = pod.Name
			p.Net = s.Name
			p.Interface = s.Interface
			p.Address = append(s.IPs)
			p.Mac = s.Mac
			outputs = append(outputs, p)
			if o.outputFormat == "text" {
				rows = append(rows, metav1.TableRow{
					Cells: p.ConvertRow(),
				})
			} else {
				rows = append(rows, metav1.TableRow{
					Cells: p.ConvertWideRow(),
				})
			}
		}
	}

	switch o.outputFormat {
	case "json":
		err = o.ShowDefaultOutputJSON(outputs)
	case "text":
		err = o.ShowDefaultOutputText(rows)
	case "wide":
		err = o.ShowWideOutputText(rows)
	default:
		err = o.ShowDefaultOutputText(rows)
	}

	return err
}

/*
// TBD for DeviceInfo of network status
func (o *PodnetOptions) ShowPCIOutput() error {
	fmt.Printf("TBD:PCI")
	return nil
}

func (o *PodnetOptions) ShowVdpaOutput() error {
	fmt.Printf("TBD:vdpa")
	return nil
}

func (o *PodnetOptions) ShowVhostUserOutput() error {
	fmt.Printf("TBD:VhostUser")
	return nil
}

func (o *PodnetOptions) ShowMemifOutput() error {
	fmt.Printf("TBD:Memif")
	return nil
}
*/
