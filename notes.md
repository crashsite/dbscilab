



# Requirements
1. github account
2. public ssh key file registered in github 


# Presentation


## Glossary
git 
workflow
github.com
    account
    organization
workspace
fork
application repo



## Revision control exists for two reasons.
1. The first is to help the act of writing code. You need to sync changes with teammates, and regularly back up your work. Emailing zip files doesn’t scale.
2. The second reason is configuration management. This includes managing parallel lines of development, such as working on the next major version while applying the occasional bug fix to the existing version in production. Configuration management is also used to figure out when exactly something changed, an invaluable tool in diagnosing bugs.


## Simple Workflow
1. Create a branch off master
2. Make changes to the branch
3. Merge branch into master


## Collabrative Workflow
1.  Fork application repo
2.  Create a reference to the application repo
3.  Create a branch off forked master
4.  Make changes to branch 
5.  Push new branch to fork
6.  Create a pull request to merge new branch from fork into application repo
7.  Merge pull request
8.  Sync up fork master


## Setup workshop
1. Go to the crashite organization on github
        https://github.com/orgs/crashsite/dashboard

2. Fork the dbscilab into your own github account

3. Clone the freshed forked repo to your workspace machine or vm
        cd ~/playground/go/src/github.com/letterj
        git clone git@github.com:letterj/dbscilab.git

4. Create Reference "upstream" for the application repo locally
        git remote add upstream git@github.com:crashsite/dbscilab.git

5. Run origin/master refresh
        git branch  # verify you are on the master branch
        git fetch upstream
        git merge upstream/master
        git push origin master


## Change hello to greetings
1.  Run origin/master refresh 
        git branch  # verify you are on the "master" branch
        git fetch upstream
        git merge upstream/master
        git push origin master

2. Create a branch for the change 
        git checkout -b greetings
        git branch   # verify you on the "greetings" branch
        sed -i /hello/greetings/g main.go
        git status
        git add main.go
        git commit -m "Changed hello to greetings"
        git log 
        git push origin greetings

3. Go to github.com and create the pull request

4. Merge the pull request


## Refresh your workspace to keep it current
1. Make sure you are on the master branch
        git checkout master

2. Run origin/master refresh
        git branch  # verify you are on the "master" branch
        git fetch upstream
        git merge upstream/master
        git push origin master

3. Clean up old branches
        git push origin :greetings  # Deletes branch from your fork
        git branch -D greetings     # Deletes branch from your workspace


