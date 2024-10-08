package cmd

import (
	"fmt"

	"github.com/fatih/color"
)

func greet() {
	nameASCIIFirst := `
 ####                                       ###        
 ######                                   #######      
 ########                       ######   #########     
 ##########                    ########  ###   ##      
 ####  #####                   ###  ###  ###           
 ####   #####                  ###       ###           
 ####    #####                 ###       ###           
 ####     #####                ###       ###           
 ####     #####                ###       ###           
 ####      #####               ###       ###          
 ####      #####               ###       ###       ### 
 ####      #####  ###   ####   ###       ###    ###### 
 ####      #####  ###   ####   ###       ### #######  
 ####      #####  ###   ####   ###       ########      
 ####       ####  ###   ####   ###      ######         
 ####       ####  ###   ####   ###   ########          
 ####       ####  #### #####   ##############          
 ####      #####  #########    #######   ####          
 ####      #####   ########  ######       ###          
 ####      ####      ###  ########        ###          
 ####     #####        ####### ###        ###          
 ####    #####       ######    ###   ###  ###          
 ####   #####      #####    #   ##  #### ####          
 #### ######      ###     ####  ##  ########           
 #########                ########   ######            
 #######                   #######                     
 #####                      ####                       
	`
	color.New(color.FgBlue).Print(nameASCIIFirst)
	fmt.Println()

}
