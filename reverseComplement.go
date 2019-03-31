package main
import (
	"fmt"
	"strings"
	"bufio"
	"os")

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
   
			
		if !(strings.HasPrefix(scanner.Text(), ">")) {
	


			fmt.Println(reverseComplement(strings.ToUpper(scanner.Text())))		}
	}
	
	
}


func reverseComplement(seq string) string {
	complement := map[string]string{"A": "T", "C": "G", "G": "C", "T": "A", "N": "N"}
	Cseq := ""
	for i := 0; i < len(seq); i++ {
		base := string(seq[i])
		Cseq = complement[base] + Cseq
	}

	return Cseq
}
