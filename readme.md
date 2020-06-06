# SonarQube and gosec integrations

Sample code for integrating with sonarqube and gosec

## Configuration

1. Run sonarqube in docker

```bash
docker run -d --name sonarqube -p 9000:9000 sonarqube
```

2. Log into <http://localhost:9000/sessions/new> with user and password `admin`
3. Click on the `+` and then `Create new project` set the project key to `your-project-name`.
4. On the new page generate a token with name `project`
5. Copy the token generated and replace the existing one for property `sonar.login` on the file `sonar-project.properties`

6. Install sonar-scaner from here <https://docs.sonarqube.org/latest/analysis/scan/sonarscanner/>
7. Run all steps for getting the reports

```bash
# install gosec tool
go get github.com/securego/gosec/cmd/gosec

# generate coverage file
go test -short -coverprofile=./cov.out ./...

# generate gosec report in sonarqube format
gosec -fmt=sonarqube -out report.json ./...

# run sonar-scanner
# sonnar-scaner relies on a file named sonar-project.properties by default
sonar-scanner
```

8. Access <http://localhost:9000/dashboard?id=your-project-name>
