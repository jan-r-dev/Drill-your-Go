1. As the first step, save the original `config.yaml` file as a backup in a `~/backups/${timestamp}/` folder before tampering with it.
2. Ingest the attached `config.yaml` file. You're expected to use a library outside of stdlib to interact with it.
3. Map it onto a struct that includes all fields.
4. Alter:
    * system_images.etcd to a different version of `rancher/coreos-etcd`
    * Internal IP of example.com to some other IP address.
    * The port of the control plane to some other port.
5. Save the modified `config.yaml` file in the place of the original.