# go-env

This is a go-env application which prints all environment variables defined in the container

To deploy this application, execute the following commands:

  1. Clone repo:

    ```
    $ git clone https://github.com/damianjaniszewski/go-env
    $ cd go-env
    ```

  2. Deploy to Kubernetes

    ```
    $ kubectl apply -f k8s-go-env-0.0.10.yaml
    $ kubectl apply -f k8s-go-env-0.1.2.yaml
    ```
  
  3. Deploy to Kubernetes with Helm

    ```
    $ helm install --name go-env helm/go-env -f helm-values-go-env-0.0.10.yaml
    $ helm install --name go-env-canary helm/go-env -f helm-values-go-env-0.1.2.yaml
    ```

  4. Deploy to CF

    ```
    $ cf push --stack cflinuxfs2 -n
    ```
