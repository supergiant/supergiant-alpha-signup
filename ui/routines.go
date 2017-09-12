package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func ConfigEnv(a *App, i Invite) {
	//for row.Next() {

	// TODO: call function to launch helm chart
	//}
	//}
	customer := i.URL
	// customer := "67tlog534wydz5bg"
	helmJSON := []byte(`{
  "chart_name": "supergiant",
  "chart_version": "0.1.0",
  "config": {
    "api": {
      "enabled": true,
      "image": {
        "pullPolicy": "Always",
        "repository": "supergiant/supergiant-api",
        "tag": "unstable-linux-x64"
      },
      "name": "supergiant-api",
      "resources": {},
      "service": {
        "externalPort": 80,
        "internalPort": 8080
      },
      "support": {
        "enabled": true,
        "password": "` + a.SupportPass + `"
      }
    },
    "ingress": {
      "annotations": {
        "traefik.frontend.rule.type": "PathPrefixStrip"
      },
      "enabled": true,
      "name": "supergiant"
    },
    "persistence": {
      "accessMode": "ReadWriteOnce",
      "enabled": true,
      "size": "8Gi",
      "storageClass": "generic"
    },
    "ui": {
      "enabled": true,
      "image": {
        "pullPolicy": "Always",
        "repository": "supergiant/supergiant-ui",
        "tag": "unstable-linux-x64"
      },
      "name": "supergiant-ui",
      "replicaCount": 1,
      "resources": {},
      "service": {
        "externalPort": 80,
        "internalPort": 3001
      }
    },
    "uniqueurl": "` + customer + `"
  },
  "kube_name": "sgalpha1",
  "name": "` + customer + `",
  "repo_name": "supergiant",
  "namespace": "` + customer + `"
}`)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	// _ = helmJSON
	_, err := client.Get("https://admin.alpha.supergiant.io/api/v0/helm_releases")
	if err != nil {
		log.Error(err)
		log.Error("Failed to launch")
	}

	log.Debug(string(helmJSON))

	req, err := http.NewRequest("POST", "https://admin.alpha.supergiant.io/api/v0/helm_releases", bytes.NewBuffer(helmJSON))
	if err != nil {
		log.Error(err)
		log.Error("Failed to launch")
	}
	req.Header.Add("Authorization", `SGAPI token="`+a.APIToken+`"`)
	req.Header.Add("Content-Type", `application/json`)
	log.Debug(req)
	log.Debug("------")
	log.Debug(req.Body)
	resp, err := client.Do(req)
	log.Debug(resp)
	bs, _ := ioutil.ReadAll(resp.Body)
	log.Debug(string(bs))
	if err != nil {
		log.Error(err)
		log.Error("Failed to launch")
	}

	var respJSON HelmRelease
	json.Unmarshal(bs, &respJSON)
	log.Debug("Create Release response JSON:")
	log.Debug(respJSON)
	releaseID := respJSON.ID
	log.Debug("Release ID")
	log.Debug(releaseID)
	//
	// releaseID := 65
	userPass := RandomString(16)
	waitErr := WaitFor("Helm Deployment", 5*time.Minute, 3*time.Second, func() (bool, error) {
		_, err = client.Get("https://admin.alpha.supergiant.io/api/v0/helm_releases/" + strconv.Itoa(releaseID))
		if err != nil {
			log.Error(err)
			log.Error("Failed to launch")
			return false, err
		}

		log.Debug(releaseID)
		req, err := http.NewRequest("GET", "https://admin.alpha.supergiant.io/api/v0/helm_releases/"+strconv.Itoa(releaseID), nil)
		log.Debug(req.URL)
		req.Header.Add("Authorization", `SGAPI token="`+a.APIToken+`"`)
		req.Header.Add("Content-Type", `application/json`)
		if err != nil {
			log.Error(err)
			log.Error("Failed to launch")
			return false, nil
		}
		resp, err := client.Do(req)
		log.Debug("Helm verify results")
		log.Debug(err)
		if err != nil {
			return false, err
		}

		bs, _ := ioutil.ReadAll(resp.Body)
		log.Debug(string(bs))

		if resp.StatusCode != 200 {
			return false, nil
		}

		var respJSON HelmRelease
		json.Unmarshal(bs, &respJSON)
		log.Debug("Release ID response JSON")
		log.Debug(respJSON)
		releaseStatus := respJSON.Status.Description

		if releaseStatus == "deploying" {
			return false, nil
		}

		_, err = client.Get("https://alpha.supergiant.io/" + customer + "/server/api/v0/sessions")
		if err != nil {
			log.Error(err)
			return false, nil
		}
		loginJSON := []byte(`{"user":{"username":"support", "password":"` + a.SupportPass + `"}}`)
		req, err = http.NewRequest("POST", "https://alpha.supergiant.io/"+customer+"/server/api/v0/sessions", bytes.NewBuffer(loginJSON))
		log.Debug(req.URL)
		req.Header.Add("Content-Type", `application/json`)
		if err != nil {
			log.Fatal(err)
		}

		resp, err = client.Do(req)
		if err != nil {
			log.Error(err)
			return false, nil
		}

		bs, _ = ioutil.ReadAll(resp.Body)
		log.Debug(string(bs))
		log.Debug(resp.StatusCode)
		if resp.StatusCode != 201 {
			log.Debug("Non 201 from session")
			return false, nil
		}

		var sesJSON Session
		err = json.Unmarshal(bs, &sesJSON)
		if err != nil {
			log.Error(err)
			return false, nil
		}

		log.Debug(sesJSON)

		userJSON := []byte(`{"username":"superadmin","password":"` + userPass + `","role":"admin"}`)
		_, err = client.Get("https://alpha.supergiant.io/" + customer + "/server/api/v0/users")
		if err != nil {
			log.Error(err)
			return false, nil
		}
		req, err = http.NewRequest("POST", "https://alpha.supergiant.io/"+customer+"/server/api/v0/users", bytes.NewBuffer(userJSON))
		log.Debug(req.URL)
		req.Header.Add("Authorization", `SGAPI session="`+sesJSON.ID+`"`)
		req.Header.Add("Content-Type", `application/json`)
		if err != nil {
			log.Fatal(err)
		}
		resp, err = client.Do(req)
		if err != nil {
			log.Error(err)
			return false, nil
		}

		log.Debug(resp.Header.Get("Content-Type"))
		bs, _ = ioutil.ReadAll(resp.Body)
		log.Debug(string(bs))
		log.Debug(resp.StatusCode)
		if resp.StatusCode != 201 {
			log.Error("Failed to create superadmin")
			return false, nil
		}
		return resp.StatusCode == 201, nil
	})
	// Done waiting
	if waitErr != nil {
		a.sendEmail("jordan@supergiant", "FAILED PROVISION", i.Invite+" - "+i.Email+"-"+customer)
		log.Fatal(waitErr)
	}

	emailBody := `Welcome to the SuperGiant Alpha.

Your environment has been configured. You can log in at https://alpha.supergiant.io/` + customer + `/ui/
with the following credentials:

username: superadmin
password: ` + userPass + `

Please change your password once logged in.`
	a.sendEmail(i.Email, "Welcome to the SuperGiant Alpha", emailBody)
}
