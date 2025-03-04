# Go Benchmarker

Go benchmarker offers a highlevel overview of your code's speed. Go benchmarker, firstly, creates a benchmark report, which is the base line report of your code's speed. After this you can create new reports after you have made changes to your code to, and go benchmarker will run a comparison between the old and new report, highlighting speed changes in a simple cli report.

## Usage
Firstly, you will need to install go benchmarker:
```shell
go install github.com/alexroden/go-benchmarker@latest
```

Once you have installed this you can then start to create reports. In the project that you want to run a report for, run:
```shell
 go-benchmarker create
```
This will create your benchmark report, the after you have made any changes to your code you can then run this command again to create a new benchmark report.

Once you have an old and new report you can then run:
```shell
go-benchmarker compare

```
The result of this will be a comparison report, which highlights speed changes in your code:
```shell
Benchmark Time Differences:
---------------------------------
client       | ▼ Decreased by 2.011s
request      | ▲ Increased by 1.581s
service/example | No change
```
Onc you are happy with the speed changes you can then accept the changes, by running:
```shell
go-benchmarker accept
```
This will then overwrite benchmark report with the new report

## License
[MIT](https://choosealicense.com/licenses/mit/)