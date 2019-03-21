# Simplified Viper Interface - A small and simple package for creating vipers (github.com/spf13/viper).



Vipers are normally created this way:

```go
my_Viper:= viper.New ()
my_Viper.SetConfigName ("file_name_without_extension")
my_Viper.SetConfigType ("yaml")
my_Viper.AddConfigPath (".")
```



However, if a file path (such as "/home/user/file.yaml") is provided, instead of the file name and its search directory, then some manual parsing will have to be done, to extract its name and path. This package simply helps with the parsing.



```go
import (

	"github.com/spf13/viper"

	viper_Interface "github.com/qamarian-inc/viper"

)

...

my_Viper, err := viper_Interface.New_Viper ("/home/user/conf.yml", "yaml")
```

