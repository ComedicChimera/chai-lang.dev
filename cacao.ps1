if ($args.Length -lt 1) {
    echo "Too few parameters supplied."
} else {
    $command = $args[0]

    switch ($command) {
        "setup" {
            npm install
            pip install "-r" requirements.txt
        }
        "rundev" {
            python chaisite/manage.py runserver localhost:8000
        }
        "build" {
            npm run build
        }
        "create-app" {
            if ($args.length -ne 2) {
                echo "Missing 'app-name' parameter"
            } else {
                cd chaisite
                python manage.py startapp $args[1]
                cd ../
            }
        }
    }
}

