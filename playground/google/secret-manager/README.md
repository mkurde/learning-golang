# Google Secret Manager

The [Secret Manager](https://cloud.google.com/secret-manager)  is a secure and
convenient storage system for API keys, passwords, certificates, and other
sensitive data. Secret Manager provides a central place and single source of
truth to manage, access, and audit secrets across Google Cloud.

## Documentation & Examples

- [gcloud beta secrets Documentation](https://cloud.google.com/sdk/gcloud/reference/beta/secrets)
  - [Create](https://cloud.google.com/sdk/gcloud/reference/beta/secrets/create)
  - [Version add / Update](https://cloud.google.com/sdk/gcloud/reference/beta/secrets/versions/add)
  - [Version access / Read](https://cloud.google.com/sdk/gcloud/reference/beta/secrets/versions/access)
- [Secret Manager Client Libraries](https://cloud.google.com/secret-manager/docs/reference/libraries)
- [Go Examples](https://github.com/GoogleCloudPlatform/golang-samples/tree/master/secretmanager)
- [Go doc secretmanager](https://pkg.go.dev/cloud.google.com/go/secretmanager/apiv1?tab=doc#pkg-overview)
-

## Blogs

- [Secret Manager: Improve Cloud Run security without changing the code](https://medium.com/google-cloud/secret-manager-improve-cloud-run-security-without-changing-the-code-634f60c541e6)

## Local development

- Create service account
- Export service account in current env `export GOOGLE_APPLICATION_CREDENTIALS=/path/to/sa.json`
- Create secret
- Update secret
- Read secret

```bash
# Create
echo "secret_value" | gcloud beta secrets create \
  --data-file=- --replication-policy=automatic my-secret-json --project=${PROJECT_NAME}

## Update
echo "secret_value updated" | gcloud beta secrets versions add \
 my-secret-json --data-file=- --project=${PROJECT_NAME}

#Read
gcloud beta secrets versions access latest --secret=my-secret-json --project=${PROJECT_NAME}
```
