# k8sConfigTest
Simple Go app to test external configuration in Kubernetes.  The API returns a json response with the value indicating where
the one configured environment variable is retrieved from.

For this API possible values, retrieved in order of precedence

 - environment
 - configuration-file
 - default

This project relies on the functionality of [viper](https://github.com/spf13/viper)

## route

http://host:80/

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
