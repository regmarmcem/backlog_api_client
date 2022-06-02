# What is this?
You can create backlog tickets from formatted excel file.

# How to use

```bash
export ENDPOINT_URL=https://<your-backlog-endpoint> # without tracing slash  
export API_KEY=<your-api-key>  
go run <excel_file_path>   
```

# Excel file format
Example.  

|projectId|summary    |priorityId|issueTypeId|assigneeId|(option1)|(option2..)|
|---------|-----------|----------|-----------|----------|---------|-----------|
|1        |urgent task|1         |23113      |11111     |foo      |bar        |

If you want to add additional options, append them to last column.  
You can check available options and valid value for each options in the following link.  
https://developer.nulab.com/ja/docs/backlog/api/2/add-issue/#
