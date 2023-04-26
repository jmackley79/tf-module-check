package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"
    "regexp"
    "strings"
)

type Module struct {
    Name    string
    Source  string
    Version string
}

type Release struct {
    Name string `json:"tag_name"`
}

func main() {
    dir, err := os.Getwd()
    if err != nil {
        fmt.Printf("Error getting current working directory: %s\n", err)
        return
    }

    files, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Printf("Error reading directory %s: %s\n", dir, err)
        return
    }

    for _, file := range files {
        if !file.IsDir() && filepath.Ext(file.Name()) == ".tf" {
            filePath := filepath.Join(dir, file.Name())

            content, err := ioutil.ReadFile(filePath)
            if err != nil {
                fmt.Printf("Error reading file %s: %s\n", filePath, err)
                continue
            }

            // Regular expression to match module and version information
            re := regexp.MustCompile(`module\s+"(.+)"\s+{\s+source\s+=\s+"(.+)"\s+version\s+=\s+"(.+)"`)

            matches := re.FindAllStringSubmatch(string(content), -1)

            if matches != nil {
                for _, match := range matches {
                    module := Module{
                        Name:    match[1],
                        Source:  match[2],
                        Version: match[3],
                    }

                    // Parse the GitHub repository and module name from the source URL
                    parts := strings.Split(module.Source, "/")
                    repoOwner := parts[0]
                    repoName := parts[1]

                    apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/terraform-aws-%s/releases/latest", repoOwner, repoName)

                    req, err := http.NewRequest("GET", apiUrl, nil)
                    if err != nil {
                        fmt.Printf("Error creating HTTP request: %s\n", err)
                        continue
                    }

                    client := &http.Client{}
                    resp, err := client.Do(req)
                    if err != nil {
                        fmt.Printf("Error sending HTTP request: %s\n", err)
                        continue
                    }

                    defer resp.Body.Close()

                    if resp.StatusCode != http.StatusOK {
                        fmt.Printf("Error: API returned status %d\n", resp.StatusCode)
                        continue
                    }

                    var release Release
                    if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
                        fmt.Printf("Error decoding API response: %s\n", err)
                        continue
                    }

                    latestVersion := strings.TrimPrefix(release.Name, "v")
                    if latestVersion != module.Version {
                        fmt.Printf("Alert: A newer version (%s) of module %s is available in file %s.\n", latestVersion, module.Name, filePath)
			// Printing URL to latest rev.
			fmt.Printf("https://github.com/%s/terraform-aws-%s/releases/latest \n", repoOwner, repoName)
                    }
                }
            }
        }
    }
}
