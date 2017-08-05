package laterutil

import "time"

func TimeToString(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

func TimeFromString(s string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil, err
	}

	tUTC := t.UTC()
	return &tUTC, nil
}
