package cmd

import (
	"fmt"
	"github.com/ubpat16/task-tracker/internal/fs"
	"os"
	"sort"
	"text/tabwriter"
)

func ListTask(f *fs.FileSystem, listType fs.ListType) {
	tasks := f.Tasks.List(listType)
	// 1. Initialize a tabwriter
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// 2. Print a nice header
	fmt.Fprintln(w, "ID\tTASK\tStatus")
	fmt.Fprintln(w, "--\t----\t----")

	// 3. Extract and sort keys so the map prints in numerical order
	var keys []int
	for k := range tasks {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// 4. Print each row separated by a tab character (\t)
	for k := range tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\n", tasks[k].ID, tasks[k].Description, tasks[k].Status)
	}

	// 5. Flush to apply the column width calculations and print
	w.Flush()
}
