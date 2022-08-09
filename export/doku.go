package export

import (
    "fmt"
    "os"
    "os/exec"
    "sort"
    "strings"

    "github.com/nboughton/swnt/content/format"
    "github.com/nboughton/swnt/content/sector"
)

// Doku represents the Exporter for Doku projects
type Doku struct {
    Name  string
    Stars *sector.Stars
}


// Write satisfies the Setup requirement of the Exporter interface
func (d *Doku) Write() error {
    fmt.Println("Exporting as dokuwiki site...")
    wdir, _ := os.Getwd()

    dokuDir := "doku"
    if err := os.Mkdir(dokuDir, dirPerm); err != nil {
        return err
    }

    if err := os.Chdir(dokuDir); err != nil {
        return err
    }

    o, err := exec.Command("git", "init").CombinedOutput()
    if err != nil {
        return err
    }
    fmt.Print(string(o))

    fmt.Println("Creating Stars dir...")
    //Create directory name by lowercasing the sector name
    //and replacing spaces with underscores
    starsDir := strings.ToLower(strings.Replace(d.Name, " ", "_", -1))
    if err := os.Mkdir(starsDir, dirPerm); err != nil {
        return err
    }
    if err := os.Chdir(starsDir); err != nil {
        return err
    }

    //Create array to hold 
    var starNames []string

    //Create formatted line
    var line string

    fmt.Println("Populating Stars dir...")
    for _, star := range d.Stars.Systems {
        // Create article stub
        starFile := strings.ToLower(strings.Replace(star.Name, " ", "_", -1)) + ".txt"
        fmt.Println("Populating file for ", starFile, "...")
        f, err := os.Create(starFile)
        if err != nil {
            return err
        }

        //Save star names so we can create an index
        starNames = append(starNames, star.Name)

        line = format.Header(format.DOKUWIKI, 1, star.Name)
        if _, err := f.Write([]byte(line)); err != nil {
            return err
        }
        line = format.Header(format.DOKUWIKI, 2, "Description")
        if _, err := f.Write([]byte(line)); err != nil {
            return err
        }
        if _, err := f.Write([]byte(star.Format(format.DOKUWIKI))); err != nil {
            return err
        }

        f.Close()
    }

    //Create star index file
    sort.Strings(starNames)
    indexFile := fmt.Sprintf("stars.txt")
    f, err := os.Create(indexFile)
    if err != nil {
        return err
    }

        line = format.Header(format.DOKUWIKI, 2, "Description")
    //Write header
    headers := [...]string{
        format.Header(format.DOKUWIKI, 3, d.Name),
        "\n",
        format.Header(format.DOKUWIKI, 4, "Description"),
        "\n",
        format.Header(format.DOKUWIKI, 4, "Sector Map"),
        fmt.Sprintf("[[map|Star Map]]\n"),
        "\n",
        format.Header(format.DOKUWIKI, 4, "Stars"),
        "\n",
    }
    for _, h := range headers {
        f.WriteString(h)
    }

    //Write each Star and link
    for _, name := range starNames {
        //link is same as file name but without the .txt
        link := strings.ToLower(strings.Replace(name, " ", "_", -1))
        line := fmt.Sprintf("| [[%s|%s]] | |\n", link, name)
        f.WriteString(line)
    }
    f.Close()
    

    // Print hexmap to map.txt
    f, err = os.Create("map.txt")
    if err != nil {
        return err
    }

    line = "=====" + d.Name + "=====\n\n<code>\n" + Hexmap(d.Stars, false, false) + "\n</code>"
    _, err = f.WriteString(line)
    if err != nil {
        return err
    }
    f.Close()

    fmt.Println("Sector written to ", wdir)

    return os.Chdir(wdir)
}
