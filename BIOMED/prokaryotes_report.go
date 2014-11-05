// This package is to retreive prokaryotes information
// from the folder of GENOME_REPORTS in NCBI ftp.
package report

import (
	"bufio"
	"io"
	"io/ioutil"
	"strings"
)

type Report struct {
	Name      string   // name of the species
	Reference []string // reference strains of the species
	Strains   []Strain // strains of the species
}

type Strain struct {
	Id          string // bioproject uid
	TaxId       string // tax id
	Name        string // name of the strain
	Genome      string // NC name
	Path        string // dir path in the NCBI FTP
	GeneticCode string // genetic table id
}

// read prokaryotes.txt
// returns a map[bioproject uid][Strain]
func ReadProkaryotes(f io.Reader) map[string]Strain {
	r := bufio.NewReader(f)
	strains := make(map[string]Strain)
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}

		if l[0] == '#' {
			continue
		}

		fields := strings.Split(l, "\t")
		if strings.Contains(l, "Complete") {
			name := fields[0]
			taxid := fields[1]
			bioid := fields[3]
			refid := strings.Split(fields[8], ".")[0]
			if strings.Contains(refid, "NC") {
				s := Strain{
					Name:   name,
					Id:     bioid,
					TaxId:  taxid,
					Genome: refid,
				}
				strains[bioid] = s
			}
		}
	}

	return strains
}

// read prok_reference_genomes.txt
// returns a map[Species name][Chromosome RefSeq]
func ReadProkReferenceGenomes(f io.Reader) map[string][]string {
	r := bufio.NewReader(f)
	species := make(map[string][]string)
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}

		if l[0] == '#' {
			continue
		}

		fields := strings.Split(l, "\t")
		name := fields[0]
		genome := fields[3]

		species[name] = append(species[name], genome)
	}

	return species
}

// read dir to get the path name
// returns a map[bioproject uid][path]
func ReadPath(dirname string) map[string]string {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}

	pathmap := make(map[string]string)
	for _, f := range files {
		if strings.Contains(f.Name(), "uid") {
			uid := strings.Split(f.Name(), "uid")[1]
			pathmap[uid] = f.Name()
		}
	}

	return pathmap
}
