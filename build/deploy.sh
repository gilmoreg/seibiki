set -e # Abort script at first error
set -u # Disallow unset variables
# Only run when not part of a pull request and on the master branch
if [ $TRAVIS_PULL_REQUEST != "false" -o $TRAVIS_BRANCH != "master" ]
then
    echo "Skipping deployment on branch=$TRAVIS_BRANCH, PR=$TRAVIS_PULL_REQUEST"
    exit 0;
fi
wget -qO- https://toolbelt.heroku.com/install-ubuntu.sh | sh
docker login -u=$DOCKER_USERNAME -p=$HEROKU_API_KEY registry.heroku.com
docker build --file ./build/Dockerfile . 
heroku container:push web --app seibiki
heroku container:release web --app seibiki
