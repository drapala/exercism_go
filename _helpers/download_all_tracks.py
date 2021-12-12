import requests
import re
import os

# Replace this with your language
LANGUAGE = "go"

# Get the html from the exercism page
pageURL = 'https://exercism.org/tracks/' + LANGUAGE + '/exercises'
request = requests.get(pageURL)
pageHTML = request.content.decode('utf-8')

# Each exercise is in the form "/tracks/go/"exerciseName&
regexString = r'tracks/' + LANGUAGE + '/exercises/(.*?)&'
allMatches = re.findall(regexString, pageHTML)
numDownloads = 0
for exercise in allMatches:
    downloadCommand = 'exercism download --exercise=' + exercise + ' --track=' + LANGUAGE
    os.system(downloadCommand)
    numDownloads += 1

print("Total downloads: ", numDownloads)