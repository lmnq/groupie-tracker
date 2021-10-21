package apiserver

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func relationsToArtists() {
	for i := range Artists {
		Artists[i].DatesLocations = Relations.Index[i].DatesLocations
	}
}

// Search ..
func Search(input, category string) ([]Artist, error) {
	var result []Artist
	switch category {
	case "artist":
		for i, k := range Artists {
			if k.Name == input {
				result = append(result, Artists[i])
			}
		}
		return result, nil
	case "creationDate":
		input, err := strconv.Atoi(input)
		if err != nil {
			return result, errors.New("Bad Request")
		}
		for i, k := range Artists {
			if k.CreationDate == input {
				result = append(result, Artists[i])
			}
		}
		return result, nil
	case "firstAlbum":
		for i, k := range Artists {
			if k.FirstAlbum == input {
				result = append(result, Artists[i])
			}
		}
		return result, nil
	case "location":
		for i, k := range Artists {
			for v := range k.DatesLocations {
				if v == input {
					result = append(result, Artists[i])
					break
				}
			}
		}
		return result, nil
	case "member":
		for i, k := range Artists {
			for _, v := range k.Members {
				if v == input {
					result = append(result, Artists[i])
					break
				}
			}
		}
		return result, nil
	}
	return result, errors.New("Bad Request")
}

// Filter ..
func Filter(r *http.Request) ([]Artist, error) {
	var result []Artist
	values := make(map[string]int)
	// r.ParseForm()
	forms := []string{"cdfrom", "cdto", "fafrom", "fato", "nomfrom", "nomto"}
	locations := r.Form["location"]
	for _, v := range forms {
		value := r.FormValue(v)
		if value == "" {
			if strings.Contains(v, "from") {
				values[v] = 0
			} else {
				values[v] = 2022
			}
			continue
		}
		ans, err := strconv.Atoi(value)
		if err != nil {
			return result, errors.New("bad request")
		}
		values[v] = ans
	}

	for i, k := range Artists {
		ctr := 0
		if k.CreationDate >= values["cdfrom"] && k.CreationDate <= values["cdto"] {
			ctr++
		}
		if len(k.Members) >= values["nomfrom"] && len(k.Members) <= values["nomto"] {
			ctr++
		}
		val := strings.Split(k.FirstAlbum, "-")
		fad, _ := strconv.Atoi(val[len(val)-1])
		if fad >= values["fafrom"] && fad <= values["fato"] {
			ctr++
		}
		if len(locations) == 0 {
			ctr++
		} else {
			for v := range k.DatesLocations {
				found := false
				for _, h := range locations {
					if h == v {
						found = true
					}
				}
				if found {
					ctr++
					break
				}
			}
		}
		if ctr >= 4 {
			result = append(result, Artists[i])
		}
	}
	return result, nil
}

// UniqLocations ..
func UniqLocations() []string {
	var result []string
	for _, v := range Artists {
		for k := range v.DatesLocations {
			found := false
			for _, h := range result {
				if k == h {
					found = true
					break
				}
			}
			if !found {
				result = append(result, k)
			}
		}
	}
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			x, z := strings.Split(result[i], "-"), strings.Split(result[j], "-")
			if x[len(x)-1] > z[len(z)-1] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result
}
