package main

import (
    "log"

    sdk "github.com/gaia-pipeline/gosdk"
)

// This is one job. Add more if you want.
func DoSomethingAwesome() error {
    log.Println("This output will be streamed back to gaia and will be displayed in the pipeline logs.")

    // An error occured? Return it back so gaia knows that this job failed.
    return nil
}

func main() {
    jobs := sdk.Jobs{
        sdk.Job{
            Handler:     DoSomethingAwesome,
            Title:       "DoSomethingAwesome",
            Description: "This job does something awesome.",

            // Increase the priority if this job should be executed later than other jobs.
            Priority: 0,
        },
    }

    // Serve
    if err := sdk.Serve(jobs); err != nil {
        panic(err)
    }
}
