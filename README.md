# WSO2 Dockerfiles test framework

WSO2 Dockerfiles provides Dockerfiles for many WSO2 products and allows users to build and run Docker images of these products. The WSO2 Dockerfiles project is still under development; changes are constantly being applied to the codebase by multiple developers of WSO2 community. In order to ensure the availability and robustness of WSO2 Dockerfiles, we are developing a general purpose integration test framework for Dockerfiles, which aims at ensuring that changes made to WSO2 Dockerfiles does not break existing functionaly and there are no regressions introduced.

## Build 
These instuctions are sepecific to Mac OSX, there should be similar counter part for Linux.  
+ Install Golang
  * Install using home brew ` brew install go --cross-compile-common `
  * Create default workspace folder ` mkdir $HOME/go `
  * Setup paths in your .bash_profile  
     ` export GOPATH=$HOME/go `  
     ` export PATH=$PATH:$GOPATH/bin ` 
+ Get the project source
   ` go get -d github.com/abhishek0198/test-framework `
+ Build the project,  
  * cd $GOPATH/src/github.com/abhishek0198/test-framework  
  * go build
+ Lauch the test framework using  
  ` ./test-framework `

## Running standard tests
The test framework also requires setting up project relevent to your tests. Following are the projects that you should clone:  
WSO2 Dockerfiles (https://github.com/wso2/dockerfiles)  
WSO2 Puppet Modules (https://github.com/wso2/puppet-modules)  

You will also need to download java and product specific zip files. Instructions can be found on WSO2 Dockerfiles.  

Once above setup is completed, follow following steps to run the tests:  
1. Edit TestConfigPath under <project_root>/src/common/test-config.json and rebuild the project  
2. Set dockerfileshome and carob_server_port in <project_root>/src/config/test-config.json  
2. Configure products to test along with desired provisioning in <project_root>/src/config/test-config.json  
3. Launch test using ```./main``` from bin directory  

IMPORTANT: In order to support running on Mac OSX, `carbon_server_ip` is explicity set to use docker-machine's host IP. Remove this config, if you're running on Linux

Here is a sample test config to test WSO2ESB using default and WSO2MB using puppet provisioning.

```        
{
   "testconfig":{
      "wso2_products":[
         {
            "enabled":"True",
            "name":"wso2esb",
            "version":"4.9.0",
            "provisioning_method":"default"
         },
         {
            "enabled":"False",
            "name":"wso2mb",
            "version":"3.1.0",
            "provisioning_method":"puppet",
            "platform":"default"
         }
      ],
      "output_file":"/home/abhishek/dev/test-framework/output.txt",
      "dockerfileshome":"/home/abhishek/dev/dockerfiles",
      "carbon_server_port":"9443"
   }
}
```
