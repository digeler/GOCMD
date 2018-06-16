package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"logsub/packages/nic"
	"logsub/packages/sub"
	"net/url"
	"os"
	"os/exec"
	"os/user"

	"strings"

	"github.com/fatih/color"

	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/compute/mgmt/compute"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2017-10-01/storage"
	"github.com/Azure/azure-storage-file-go/2017-07-29/azfile"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

const (
	// DefaultBaseURI is the default URI used for the service Network
	DefaultBaseURI = "https://management.azure.com"
)

var (
	ctx = context.Background()

	authorizer autorest.Authorizer
	rg         string
	global     string
)

func dooper(vmname string, rgname string, storagename string) {
	//sudo mount -t cifs //k8logcjews.file.core.windows.net/k8logs [mount point] -o vers=3.0,username=k8logcjews,password=07HVJ0a01Av4AnY1ZZsTVs5A0wiwtjSBfrKhgvzG2n3kRvrz7khjnrZtKN4/Xphu3UCjadbE6X5F2VKQ/AlcTw==,dir_mode=0777,file_mode=0777,sec=ntlmssp

	fmt.Println("Going to create the mount /mnt/forlogs/ on vm \n", vmname)

	cmd := exec.Command("az", "vm", "run-command", "invoke", "--resource-group", rgname, "--name", vmname, "--command-id", "RunShellScript", "--scripts", "mkdir /mnt/forlogs")

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Command finished with error: %v", err)
	}
	//args := []string{"vm", "run-command", "invoke", "--resource-group", rgname, "--name", vmname, "--command-id", "RunShellScript", "--scripts", "sudo", "mount", "-t cifs", "//", storagename, ".file.core.windows.net/k8logs", "/mnt/forlogs", "-o", "vers=3.0,", "username=", storagename, ",password=", global, ",dir_mode=0777", ",file_mode=0777", ",sec=ntlmssp"}
	buf := bytes.Buffer{}

	//install cifs
	buf.WriteString("az vm run-command invoke --resource-group " + rgname)
	buf.WriteString(" --name " + vmname)
	buf.WriteString(" --command-id RunShellScript --scripts ")
	buf.WriteString("'sudo apt-get update && sudo apt-get install cifs-utils'\n")
	///////////////////
	buf.WriteString("az vm run-command invoke --resource-group " + rgname)
	buf.WriteString(" --name " + vmname)
	buf.WriteString(" --command-id RunShellScript --scripts ")
	buf.WriteString("'sudo mount -t cifs //" + storagename)
	buf.WriteString(".file.core.windows.net/k8logs")
	buf.WriteString(" /mnt/forlogs/ -o vers=3.0,")
	buf.WriteString("username=" + storagename)
	buf.WriteString(",password=" + global)
	buf.WriteString(",dir_mode=0777,file_mode=0777,sec=ntlmssp'\n")
	//kubelog command **************************************
	buf.WriteString("az vm run-command invoke --resource-group " + rgname)
	buf.WriteString(" --name " + vmname)
	buf.WriteString(" --command-id RunShellScript --scripts ")
	buf.WriteString("'journalctl -u kube* --no-pager>>")
	buf.WriteString("/mnt/forlogs/" + vmname)
	buf.WriteString(".log'\n")
	//clusterlog**************************************
	buf.WriteString("az vm run-command invoke --resource-group " + rgname)
	buf.WriteString(" --name " + vmname)
	buf.WriteString(" --command-id RunShellScript --scripts ")
	buf.WriteString("'cp /var/log/azure/cluster-provision.log ")
	buf.WriteString("/mnt/forlogs/" + vmname)
	buf.WriteString(".cluster-provision.log'\n")
	//syslog//////////////////////////////////
	buf.WriteString("az vm run-command invoke --resource-group " + rgname)
	buf.WriteString(" --name " + vmname)
	buf.WriteString(" --command-id RunShellScript --scripts ")
	buf.WriteString("'mkdir /mnt/forlogs/syslog" + vmname)
	buf.WriteString(" && cp /var/log/syslog* /mnt/forlogs/syslog'" + vmname)
	buf.WriteString("\n")
	//journalall/////////////////////////////////////
	buf.WriteString("az vm run-command invoke --resource-group " + rgname)
	buf.WriteString(" --name " + vmname)
	buf.WriteString(" --command-id RunShellScript --scripts ")
	buf.WriteString("'journalctl --no-pager >>")
	buf.WriteString(" /mnt/forlogs/" + vmname)
	buf.WriteString(".journal'\n")
	//iptables/////////////////////////////////////////////////////
	buf.WriteString("az vm run-command invoke --resource-group " + rgname)
	buf.WriteString(" --name " + vmname)
	buf.WriteString(" --command-id RunShellScript --scripts ")
	buf.WriteString("'sudo iptables -S >>")
	buf.WriteString(" /mnt/forlogs/" + vmname)
	buf.WriteString(".iptable'\n")

	//packging files
	/*buf.WriteString("az vm run-command invoke --resource-group " + rgname)
	buf.WriteString(" --name " + vmname)
	buf.WriteString(" --command-id RunShellScript --scripts ")
	buf.WriteString("'tar -cvf kubelogs.tar" + vmname)
	buf.WriteString(" /mnt/forlogs/'")
	*/
	//cmd := exec.Command("az", "vm", "run-command", "invoke", "--resource-group", rgname, "--name", vmname, "--command-id", "RunShellScript", "--scripts", "mkdir /mnt/forlogs")
	mycmd := buf.String()
	fmt.Println("\nGoing to execute\n", mycmd)

	usr, _ := user.Current()

	f, err := os.Create(usr.HomeDir + "/cmd.sh")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	//var out bytes.Buffer
	//cmd.Stdout = &out
	_, err = f.WriteString(mycmd)
	if err != nil {
		fmt.Printf("Command finished with error: %v", err)
	}

	f.Sync()
	//fmt.Println(usr.HomeDir + "/cmd.sh")
	cmd = exec.Command("sh", usr.HomeDir+"/cmd.sh")

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Command finished with error(ignore is 1,2): %v\n", err)
	}
	var out bytes.Buffer
	cmd.Stdout = &out

	color.Green("Log collection on vm done you can now fetch the logs from share k8logs at storage account named " + storagename)

	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println(err)
	//}

	fmt.Printf(out.String())

}

func createstorage(ctx context.Context, storageAccountsClient storage.AccountsClient, rgname string, location string) (s storage.Account, err error) {

	saname := nic.Randname(5)
	r := strings.Replace(saname, "", "k8log", 1)
	m := strings.ToLower(r)
	fmt.Printf("creating storage account for logging name %s ", m)
	future, err := storageAccountsClient.Create(
		ctx,
		rgname,
		m,
		storage.AccountCreateParameters{
			Sku: &storage.Sku{
				Name: storage.StandardLRS},
			Kind:     storage.Storage,
			Location: to.StringPtr(location),
			AccountPropertiesCreateParameters: &storage.AccountPropertiesCreateParameters{},
		})
	if err != nil {
		return s, fmt.Errorf("cannot create storage account: %v", err)
	}

	err = future.WaitForCompletion(ctx, storageAccountsClient.Client)
	if err != nil {
		return s, fmt.Errorf("cannot get the storage account create future response: %v", err)
	}

	return future.Result(storageAccountsClient)

}

func holdkey(key string) string {
	global := key
	return global

}

func main() {
	token, subid, rgname, vnetname, loglocation, subnetName, location, sshpubkey, err := sub.Readfromauth()
	//groups client
	groupsClient := resources.NewGroupsClient(subid)
	if err != nil {
		fmt.Println(err)

	}
	groupsClient.Authorizer = token
	//print info to the customer
	fmt.Printf("\n\n")
	k := color.New(color.FgCyan, color.Bold)
	k.Printf("**********Please make sure you run az login and selected the correct sub,or else it fail **********\n\n\n")
	color.Green("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	color.Yellow("make sure the correct details below : *********************** ")
	color.Yellow("RG: %s\n", rgname)
	color.Yellow("SUBID: %s\n", subid)
	color.Yellow("VNETNAME: %s\n", vnetname)
	color.Yellow("LOGLOCATION: %s\n", loglocation)
	color.Yellow("SUBNET: %s\n", subnetName)
	color.Yellow("LOCATION: %s\n", location)
	color.Yellow("PUBKEY: %s\n", sshpubkey)

	color.Green("**************************************************************************************************************************")
	color.Red("!!!This tool relies on vm guest agent make sure agent is ok!!!!")
	color.Red("Also if the shell closes sooner as expected, you have the command at current dir name cmd.sh ,so you can run manually")
	color.Green("**************************************************************************************************************************")
	color.Green("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	//holder for sa

	//vmclient
	vmClient := compute.NewVirtualMachinesClient(subid)
	vmClient.Authorizer = token
	//storage client
	storageAccountsClient := storage.NewAccountsClient(subid)
	storageAccountsClient.Authorizer = token

	//create storage account to hold the share
	d, err1 := createstorage(ctx, storageAccountsClient, rgname, location)
	if err != nil {
		fmt.Println(err1)
	}
	n := d.Name
	gkeys, err := storageAccountsClient.ListKeys(ctx, rgname, *n)
	f := gkeys.Keys
	for _, l := range *f {

		c := holdkey(*l.Value)
		global = c
	}

	credential := azfile.NewSharedKeyCredential(*n, global)
	p := azfile.NewPipeline(credential, azfile.PipelineOptions{})
	u, _ := url.Parse(fmt.Sprintf("https://%s.file.core.windows.net", *n))
	serviceURL := azfile.NewServiceURL(*u, p)
	shareURL := serviceURL.NewShareURL("k8logs")
	_, err = shareURL.Create(ctx, azfile.Metadata{}, 0)
	if err != nil && err.(azfile.StorageError) != nil && err.(azfile.StorageError).ServiceCode() != azfile.ServiceCodeShareAlreadyExists {
		log.Fatal(err)
	}
	fmt.Println("share was created", shareURL.String())

	x, _ := vmClient.List(ctx, rgname)

	s := x.Values()

	for _, v := range s {

		l := v.Name

		dooper(*l, rgname, *n)
		//dooper(vmname string, rgname string, storagename string, key string)

		//fmt.Printf("%s\n", *l)
		//env GOOS=linux GOARCH=amd64 go build -v main.go

		//sudo mount -t cifs //k8logcjews.file.core.windows.net/k8logs [mount point] -o vers=3.0,username=k8logcjews,password=07HVJ0a01Av4AnY1ZZsTVs5A0wiwtjSBfrKhgvzG2n3kRvrz7khjnrZtKN4/Xphu3UCjadbE6X5F2VKQ/AlcTw==,dir_mode=0777,file_mode=0777,sec=ntlmssp
	}

	//cleanup phase***********************
	color.Blue("*********************************************************************")
	color.Green("!!!Please make sure you copy all the needed files from the share as we are going to delete it as part of cleanup!!!!")
	color.Green("Press 'Enter' 3 times to continue cleanup phase")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	color.Green("Press 'Enter' if you sure you want to delete the storage account")
	fmt.Println("Starting deleting storage account ", *n)
	_, err = storageAccountsClient.Delete(ctx, rgname, *n)
	if err != nil {
		color.Red("problem deleting storage \n", *n, err)
	}

}
