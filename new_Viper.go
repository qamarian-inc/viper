package viper

import (

	"github.com/spf13/viper"
	"os"
        "path"
        "path/filepath"
      	"strings"
)

func New_Viper (file_Path string, format string) (*viper.Viper, error) { // This function creates a new viper, given the path of a configuration file, and the file's format.

	// Verifying if file exists. { ...
        _, err := os.Open (file_Path)
        if err != nil {
                return nil, err
        }
        // ... }

        // Extracting file name (without extension), and directory of the file. { ...
        data_File_Name := filepath.Base (file_Path)

        data_File_Dir := strings.TrimSuffix (file_Path, data_File_Name)
        if data_File_Dir != "" {
                data_File_Dir = data_File_Dir [ : len (data_File_Dir)]
        } else {
                data_File_Dir = "."
        }

        data_File_Name_Without_Ext := strings.TrimSuffix (data_File_Name, path.Ext (data_File_Name))
        // ... }

        // Creating viper. { ...
        net_Addrses := viper.New ()
        net_Addrses.SetConfigName (data_File_Name_Without_Ext)
        net_Addrses.SetConfigType (format)
        net_Addrses.AddConfigPath (data_File_Dir)
        err2 := net_Addrses.ReadInConfig ()
        // ... }

        return net_Addrses, err2
}
