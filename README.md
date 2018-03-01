# k8sConfigTest
Simple Go app to test external configuration in Kubernetes.  The API returns a json response with the value indicating where
the one configured environment variable is retrieved from.

For this API possible values, retrieved in order of precedence

 - environment
 - configuration-file
 - default

This project relies on the functionality of [viper](https://github.com/spf13/viper)

## route

http://host:80/config

Response will be of form,
 
```
{
	'configuration': true,
	'value': 'default'
}
```

## setup for testing

### environment 
Set K8SCONFIG_VNAME = "environment"

### configuration file
Using either a JSON, TOML, YAML, HCL, or Java Properties file
with the filename of the form

add a property with the name "vName" and set it's value to "configuration-file"

```
k8sconfig.[json, toml, yaml, hcl, properties]
```

### default value
The default values is preset to "default"


## running the application

```
k8sconfig     // will default to 0.0.0.0:80

or

k8sconfig --url="0.0.0.0:yourport"
```

## Dockerfile

The docker file will perform a multistage build and allow the application to be run on port 80

```
interactive mode / with output of 'default'
docker run -p80:80 it --rm yourImageName

inject a configuration file / with output 'configuration-file'
docker run -p80:80 -it --rm -v /yourconfigfilelocation:/app/config yourImageName

inject an environment variable / with output 'environment'
docker run -p80:80 -it --rm --env K8SCONFIG_VNAME="environment" yourImageName
```