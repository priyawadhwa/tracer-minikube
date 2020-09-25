package minikube

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/pkg/errors"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
)

func start() error {
	cmd := exec.Command("minikube", "start", "--output", "json")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
		if err := processStep(m); err != nil {
			return err
		}
	}
	return nil
}

func processStep(step string) error {
	event := cloudevents.NewEvent()
	if err := json.Unmarshal([]byte(step), &event); err != nil {
		return errors.Wrap(err, "unmarshal cloud event")
	}
	m := map[string]string{}
	data := event.Data()
	if err := json.Unmarshal(data, &m); err != nil {
		return errors.Wrap(err, "unmarshal data")
	}
	stepName := m["name"]
	fmt.Println("step name is", stepName)

	// for now, assume every step takes 3 seconds
	t := 3 * time.Second

	ctx, err := tag.New(context.Background(), tag.Insert(keyMethod, stepName))
	if err != nil {
		return errors.Wrap(err, "new tag")
	}
	stats.Record(ctx, mLatencyS.M(t.Seconds()))

	return nil
}
