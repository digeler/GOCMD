This tool was created to collect logs from inside aks nodes.
how it works :
it uses the az vm run-command
here are the steps :
- First it will create storage account at the same RG as your AKS MC_
- Mount the share on the vm so we can move the logs there
-It will start to collect the logs from each node 
-Logs collected : journalctl -u kube* and cluster provision log.
-All the logs will be in the storage account.
-You Can then download the logs and share with css.

How to install and run the code :
-open cloudshel ,select the correct subscription and cd to your Clouddrive folde
-Clone the project : git clone https://github.com/digeler/GOCMD.git
-Edit the Auth file inside the GOCMD folder :

{
  "clientId": "",       -------->make sure the appid is at least owner on the rg.
  "clientSecret": "****",
  "subscriptionId": "928f4e7e-2c28-4063-a56e-6f1e6f2bb73c",
  "tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47",
  "activeDirectoryEndpointUrl": "https://login.microsoftonline.com",
  "resourceManagerEndpointUrl": "https://management.azure.com/",
  "activeDirectoryGraphResourceId": "https://graph.windows.net/",
  "sqlManagementEndpointUrl": "https://management.core.windows.net:8443/",
  "galleryEndpointUrl": "https://gallery.azure.com/",
  "managementEndpointUrl": "https://management.core.windows.net/",
  "rgname": "MC_sql1_aks1_westeurope",
  "location": "westeurope"

  }
  
  -Copy the executable to be on the same folder as auth : cp k8scollectorlinux/k8scollector .
  if all is good you can run the exec :
  ./k8scollector
  it will ask :
   please enter the full path to the auth file
   just type auth
   the code will run and at the end collect the logs and do the cleanup
   
   
  
  
  






