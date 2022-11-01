package svm_profile

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

// make reflector jangan lupa
func (profile *Profile) Write(name_step string) {

	vp := viper.New()

	// default.json
	vp.SetConfigName("default")
	vp.SetConfigType("json")
	vp.AddConfigPath("/opt/svm")
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	if name_step == "default" {
		profile.Name = "default"
	}
	// fmt.Println(vp.GetString("default"))
	profile.Uuid, _ = exec.Command("uuidgen").Output()
	vp.Set(profile.Name, map[string]interface{}{"name": profile.Name, "uuid": strings.Replace(string(profile.Uuid[:]), "\n", "", -1), "workdir": profile.Workdir, "ssh_path": profile.Ssh_path, "default_user": profile.Default_user})
	// fmt.Println(vp.GetStringMapString("default"))
	vp.WriteConfig()
	fmt.Println(name_step, "Writed")
}

// Read Profile
func (profile *Profile) Read() {
	fmt.Println("Read Profile")
	vp := viper.New()
	// Read Config
	vp.SetConfigName("default")
	vp.SetConfigType("json")
	vp.AddConfigPath("/opt/svm")
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	vp.UnmarshalKey("default", &profile)

}
