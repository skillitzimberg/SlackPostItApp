# Slack Post It App

## POST_ITS:
* Receives post-it commands from Slack.
* These commands have ACTION, CATEGORY, and NOTE
* Slash command structure: `/PostIt [ACTION] [CATEGORY] [NOTE]`
    - `/PostIt post todo get milk` will create a post-it with CATEGORY: todo; NOTE: get milk - these post-its are saved to a database.
    - `/PostIt get todo` will return all post-its with CATEGORY: todo - these post-its are retrieved from a database.

* I'm not convinced these need to be separate workflows. Use `run-if`s based on `post` or `get`.

## TOOLS
* Slack app set up
* webhook trigger
* common/template
* convert/parse_line
* database/postgres_pkg
* project values/variables

## Schema:
POSTIT:  
- Category
- Note

## ACTIONS

1. Figure out how to get a local database server going
2. Create database
3. Create post-it table
4. Find/write `post` command
5. Find/write `get` command
6. Register PostIt app with Slack
7. Create slash commands for `post` & `get`.
8. Write workflow.yaml

    
## WORKFLOW:
- convert:parse_line@1.0: 
    * parse initial Slack response on `&`; 
    * return `text` field
- convert:parse_line@1.0: 
    * parse `text` field on `=`; 
    * return command text
- convert:parse_line@1.0: 
    * parse command text on `+`; 
    * return `[ACTION]`, `[CATEGORY]`, & `[NOTE]`
- database:unknown-post@1.0:
    * run-if: `${command.action} == post && ${command.category} != null && ${command.note} != null`
    * return `Post It saved.`
- database:unknown-get@1.0:
    * run-if: `${command.action} == get && ${command.category} != null`
    * return list of requested category post-its

## PROBLEMS ENCOUNTERED
### Publishing a Step
**ERROR** `You have specified an executable for linux - amd64 but an executable was not found at /Users/scott/apptree/step_library/convert_pkg/main_linux_amd64`  
**Problem:** What is the command to build these executables? I know Ive seen it before, but I think Matthew just found it and entered it.  
**Solutions:**  
Make sure that the package that includes the steps you want to publish is included in the Makefile. Once the package is listed there run `make publish-package_name`, e.g. ` make publish-convert`.  

**ERROR** `error message: 500_service_error`  
**Problem** When attempting to run the workflow locally, the server does not start a webhook url.  
**Solutions:**  
- Make sure that the path in the `apptree run workflow` command is correct. E.g. for this command to work - `apptree run workflow -f workflows/SlackPostIt/workflow.yaml --id SlackPostIt`, you must be in the directory that contains the `workflows` directory.

### Publishing a Workflow
**ERROR**  `error message: 404_client_error`  
**Problem** The workflow & steps are successfully published. At this point the app no longer has access to my locally running, Dockerized database.  
**Solutions:**  
- install an AppTree remote engine
- host the database in the cloud and rewrite the workflow & steps as needed to make that API call. May only need to change a connection string depending on how your app is set up.