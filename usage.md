  Gocollector v1.0 
  please make sure you run the tool with the exe that suit your architecture 

  if you want to compile to linux :
  env GOOS=linux GOARCH=amd64 go build -v k8scollector.go

  for windows :
  env GOOS=windows GOARCH=amd64 go build -v -o k8scollector.exe  k8scollector.go

  for OSX :
  just compile on MAC


  
  This tool was created to collect logs from inside aks nodes.  
  how it works :  it uses the az vm run-command  here are the steps :  - First it will create storage account at the same RG as your AKS MC_  
  - Mount the share on the vm so we can move the logs there  
  -It will start to collect the logs from each node   
  -Logs collected : journalctl -u kube* and cluster provision log.  
  -All the logs will be in the storage account.  
  -You Can then download the logs and share with css.
  
   for questions you can meet me here : http://slack.thegbsguy.com

   How to install and run the code : Clone the project : git clone https://github.com/digeler/GOCMD.git -Edit the Auth file inside the GOCMD folder :

{  "clientId": "",                    -------->make sure the appid is at least owner on the rg. 
"clientSecret": "****",  "subscriptionId": "928f4e7e-2c28-4063-a56e-6f1e6f2bb73c",  
"tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47", 
"activeDirectoryEndpointUrl": "https://login.microsoftonline.com",
"resourceManagerEndpointUrl": "https://management.azure.com/",  
"activeDirectoryGraphResourceId": "https://graph.windows.net/", 
"sqlManagementEndpointUrl": "https://management.core.windows.net:8443/",
"galleryEndpointUrl": "https://gallery.azure.com/",
"managementEndpointUrl": "https://management.core.windows.net/",  
"rgname": "MC_sql1_aks1_westeurope",  "location": "westeurope"
}
     -Run az login and az account set --subscription ,to set your account 
     -Run the executable and give the path for the auth file
     
     

The MIT License (MIT)

Copyright (c) 2018 Dinor Geler 

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.  


<iframe width="560" height="315" src="https://www.youtube.com/embed/IAbXuSNHrAU" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>


