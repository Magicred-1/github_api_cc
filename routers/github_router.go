package routers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github_api/config"
	"github_api/types"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetGHAllUserRepos(c *fiber.Ctx) error {
	// Extract the username parameter from the URL
	username := c.Params("username")

	if username == "" {
		// Return an error message if the username is empty
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": "Username is required"})
	}

	// Build the GitHub API URL
	url := fmt.Sprintf("https://api.github.com/users/%s/repos?sort=updated", username)

	// Create an HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set the Accept header to application/vnd.github.v3+json
	req.Header.Add("Accept", "application/vnd.github.v3+json")

	// Send the request to GitHub API
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check the HTTP status code
	if res.StatusCode != http.StatusOK {
		// Handle errors, you can return an error message or JSON response
		return c.Status(res.StatusCode).JSON(map[string]string{"error": "Failed to fetch repositories"})
	}

	// Decode the JSON response into a slice of Repository structs
	var repos []types.Repository
	err = json.NewDecoder(res.Body).Decode(&repos)
	if err != nil {
		return err
	}

	// Capture the response body before closing it
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Pass the captured response body to writeCSVfromJSON
	err = writeCSVfromJSON(responseBody, username, "repos")
	if err != nil {
		return err
	}

	// Return the list of repositories as JSON
	return c.JSON(repos)
}

func GetGHUserRepo(c *fiber.Ctx) error {
	username := c.Params("username")
	reponame := c.Params("reponame")

	url := "https://api.github.com/repos/" + username + "/" + reponame

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return c.Status(res.StatusCode).JSON(map[string]string{"error": "Failed to fetch repository"})
	}

	// Capture the response body before closing it
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Pass the captured response body to writeCSVfromJSON
	err = writeCSVfromJSON(responseBody, username, reponame)
	if err != nil {
		return err
	}

	// Decode the JSON response into a Repository struct (assuming a single repository)
	var repo types.Repository
	err = json.Unmarshal(responseBody, &repo)
	if err != nil {
		return err
	}

	return c.JSON(repo)
}

func GetGHAllOrgRepo(c *fiber.Ctx) error {
	orgname := c.Params("orgname")

	url := "https://api.github.com/orgs/" + orgname + "/repos" + "?sort=updated"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "Bearer "+config.Config("GH_TOKEN"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return c.Status(res.StatusCode).JSON(map[string]string{"error": "Failed to fetch repositories"})
	}

	// Capture the response body before closing it
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Pass the captured response body to writeCSVfromJSON
	err = writeCSVfromJSON(responseBody, orgname, "orgs")
	if err != nil {
		return err
	}

	// Decode the JSON response into a slice of Repository structs
	var repos []types.Repository
	err = json.Unmarshal(responseBody, &repos)
	if err != nil {
		return err
	}

	return c.JSON(repos)
}

func GetGHOrgRepo(c *fiber.Ctx) error {
	orgname := c.Params("orgname")
	reponame := c.Params("reponame")

	url := "https://api.github.com/repos/" + orgname + "/" + reponame

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return c.Status(res.StatusCode).JSON(map[string]string{"error": "Failed to fetch repository"})
	}

	// Capture the response body before closing it
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Pass the captured response body to writeCSVfromJSON
	err = writeCSVfromJSON(responseBody, orgname, reponame)
	if err != nil {
		return err
	}

	// Decode the JSON response into a Repository struct (assuming a single repository)
	var repo types.Repository
	err = json.Unmarshal(responseBody, &repo)
	if err != nil {
		return err
	}

	return c.JSON(repo)
}

func DownloadRepoSource(c *fiber.Ctx) error {
	username := c.Params("username")
	reponame := c.Params("reponame")

	url := "https://github.com/" + username + "/" + reponame + "/archive/refs/heads/main.tar.gz"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "Bearer "+config.Config("GH_TOKEN"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return c.Status(res.StatusCode).JSON(res.Status)
	}

	// Create and open a file for writing the downloaded source
	if _, err := os.Stat("./public/downloads"); os.IsNotExist(err) {
		os.MkdirAll("./public/downloads", os.ModePerm)
	}

	file, err := os.Create("./public/downloads/" + reponame + ".zip")
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the downloaded source to the file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	// Return a success message
	return c.SendString("Repository source downloaded successfully")
}

func GetDownloadedRepos(c *fiber.Ctx) error {
	if _, err := os.Stat("./public/downloads"); os.IsNotExist(err) {
		os.MkdirAll("./public/downloads", os.ModePerm)
	}

	files, err := os.ReadDir("./public/downloads/")
	if err != nil {
		return err
	}

	var repos []types.Repository
	for _, file := range files {
		repos = append(repos, types.Repository{
			Name: file.Name(),
			URL:  fmt.Sprintf("http://localhost:3000/api/repos/%s/download", file.Name()),
		})
	}

	return c.JSON(repos)
}

func writeCSVfromJSON(data []byte, username string, reponame string) error {
	// Create and open a file for writing the CSV
	if _, err := os.Stat("./public/csv"); os.IsNotExist(err) {
		os.MkdirAll("./public/csv", os.ModePerm)
	}

	file, err := os.Create("./public/csv/" + username + "_" + reponame + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()

	// Initialize a CSV writer
	writer := csv.NewWriter(file)

	// Ensure the writer flushes any buffered data and writes to the file upon function completion
	defer writer.Flush()

	// Decode the JSON data into a slice of Repository structs
	var repos []types.Repository
	err = json.Unmarshal(data, &repos)
	if err != nil {
		return err
	}

	// Write the CSV header
	header := []string{"ID", "Name", "URL", "Description"}
	err = writer.Write(header)
	if err != nil {
		return err
	}

	// Write each repository entry to the CSV file
	for _, repo := range repos {
		record := []string{strconv.Itoa(repo.ID), repo.Name, repo.URL, repo.Description}
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}
