package p1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Bag struct {
	Name string
	Parent []*Bag
}

func (b *Bag) addMoreParent(newParent *Bag) {
	b.Parent = append(b.Parent, newParent)
}

func computeTotalParents(target string, childToParent map[string]*Bag) int {
	count := 0
	if parents, ok := childToParent[target]; ok {
		for _, p := range parents.Parent {
			fmt.Println(target, *&p.Name)
			count += (1 + computeTotalParents(*&p.Name, childToParent))
		}
	}
	return count
}

func driver() int {
	filename := fmt.Sprintf("p1.input")
	fp, fpe := os.Open(filename)
	if fpe != nil {
		panic(fpe)
	}
	defer fp.Close()

	childToParent := make(map[string]*Bag)

	bufReader := bufio.NewScanner(fp)
	for bufReader.Scan() {
		input := bufReader.Text()
		matcher := regexp.MustCompile(`^(.+?)\ bag[s]?\ contain\ (?:(?:[1-9]\ (.+?)\ bag[s]*)(?:,\ [1-9]\ (.+?)\ bag[s]?)*|(?:no\ other\ bags)).$`)
		matchedComponents := matcher.FindStringSubmatch(input)[1:]
		parent := matchedComponents[0]
		children := matchedComponents[1:]

		if parent == "clear gold" {
			fmt.Println(matchedComponents)
			fmt.Println(children)
		}

		parentNode := Bag{Name: parent, Parent: []*Bag{}}
		for _, child := range children {
			if existingChildNode, ok := childToParent[child]; ok {
				existingChildNode.addMoreParent(&parentNode)
			} else {
				childNode := Bag{Name: child, Parent: []*Bag{&parentNode}}
				childToParent[child] = &childNode
			}
		}
	}

	allGenerationParents := computeTotalParents("shiny gold", childToParent)
	fmt.Println(allGenerationParents)
	return 0
}
