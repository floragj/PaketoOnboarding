package main

const AppName string = "onboarding_app"

type BuildPlan struct {
	Provides []Provide `toml:"provides"`
	Requires []Require `toml:"require"`
}

type Provide struct {
	Name string `toml:"name"`
}

type Require struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
}

type Detector struct{}

func NewDetector() Detector {
	return Detector{}
}

//
// Inputs:
// platformPath:
//   path to platform sepecific extensions, not used in this example
// planPath:
//   path the build BuildPlan, where we are going to write our `provides` and `requires` statements
// appPath:
//   path to the root of the application, going to need to inspect the
//   package.json file here to figure out if our app passes detection
//
func (d Detector) DetectFunction(platformPath, planPath, appPath string) (int, error) {

	// outline of what a package.json file will look like
	var packageJSON struct {
		Name    string `json:"name"`
		Engines struct {
			NodeVersion string `json:"node"`
		} `json:"engines"`
	}

	//
	// Open our the packageJSON file and Decode it into
	// the packageJSON struct defined above
	//

	//
	// check if the packageJSON.Name field is equal to "onboarding_app"
	// if so continue,
	// if not, then we should fail detection (100 exit status)
	//

	//
	// Great at this point we know that we have an app that should
	// pass detection
	//
	// Now lets write out our BuildPlan [[provides]] and [[requires]]
	// using the BuildPlan Struct defined above
	// and the 'toml' library.
	// For an example of how to write out the toml see:
	// https://godoc.org/github.com/BurntSushi/toml#Encoder.Encode
	//

	// finally return the exit status 0!
	return 0, nil
}
