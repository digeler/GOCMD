<p>  This tool was created to collect logs from inside aks nodes.  how it works :  it uses the az vm run-command  here are the steps :  - First it will create storage account at the same RG as your AKS MC_  - Mount the share on the vm so we can move the logs there  -It will start to collect the logs from each node   -Logs collected : journalctl -u kube* and cluster provision log.  -All the logs will be in the storage account.  -You Can then download the logs and share with css.</p>

<p>How to install and run the code : -open cloudshel ,select the correct subscription and cd to your Clouddrive folde -Clone the project : git clone https://github.com/digeler/GOCMD.git -Edit the Auth file inside the GOCMD folder :</p>

<p>{  &quot;clientId&quot;: &quot;&quot;, -------->make sure the appid is at least owner on the rg.  &quot;clientSecret&quot;: &quot;****&quot;,  &quot;subscriptionId&quot;: &quot;928f4e7e-2c28-4063-a56e-6f1e6f2bb73c&quot;,  &quot;tenantId&quot;: &quot;72f988bf-86f1-41af-91ab-2d7cd011db47&quot;,  &quot;activeDirectoryEndpointUrl&quot;: &quot;https://login.microsoftonline.com&quot;,  &quot;resourceManagerEndpointUrl&quot;: &quot;https://management.azure.com/&quot;,  &quot;activeDirectoryGraphResourceId&quot;: &quot;https://graph.windows.net/&quot;,  &quot;sqlManagementEndpointUrl&quot;: &quot;https://management.core.windows.net:8443/&quot;,  &quot;galleryEndpointUrl&quot;: &quot;https://gallery.azure.com/&quot;,  &quot;managementEndpointUrl&quot;: &quot;https://management.core.windows.net/&quot;,  &quot;rgname&quot;: &quot;MC_sql1_aks1_westeurope&quot;,  &quot;location&quot;: &quot;westeurope&quot;</p>

<p> }    -Copy the executable to be on the same folder as auth : cp k8scollectorlinux/k8scollector .  if all is good you can run the exec :  ./k8scollector  it will ask :  please enter the full path to the auth file  just type auth  the code will run and at the end collect the logs and do the cleanup      to see it in action goto :    link to site: https://digeler.github.io/GOCMD/.    <iframe width=&quot;560&quot; height=&quot;315&quot; src=&quot;https://www.youtube.com/embed/IAbXuSNHrAU&quot; frameborder=&quot;0&quot; allow=&quot;autoplay; encrypted-media&quot; allowfullscreen></iframe>  </p>

<p></p>

<p></p>

<p></p>
