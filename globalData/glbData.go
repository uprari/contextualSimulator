package glb

import(
		"bufio"
        "os"
)

var(

	Peer39_data = []string{}
	GumGum_data = []string{}
)


func Init_response_data () {


	jsonFile,_ := os.Open("simResponse/peer39.json")

    defer jsonFile.Close()

	scanner := bufio.NewScanner(jsonFile) 
	for scanner.Scan() {	
		Peer39_data = append( Peer39_data,scanner.Text() )
    }

	jsonFile.Close()

	jsonFile,_ = os.Open("simResponse/gumgum.json")
	scanner = bufio.NewScanner(jsonFile)

	for scanner.Scan(){
		GumGum_data = append(GumGum_data, scanner.Text())
    }

}
