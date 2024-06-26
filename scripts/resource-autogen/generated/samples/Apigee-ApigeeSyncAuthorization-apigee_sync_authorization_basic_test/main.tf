/**
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

```hcl
resource "google_project" "project" {
  project_id      = "my-project"
  name            = "my-project"
  org_id          = "123456789"
  billing_account = "000000-0000000-0000000-000000"
}

resource "google_project_service" "apigee" {
  project = google_project.project.project_id
  service = "apigee.googleapis.com"
}

resource "google_apigee_organization" "apigee_org" {
  analytics_region   = "us-central1"
  project_id         = google_project.project.project_id

  runtime_type       = "HYBRID"
  depends_on         = [google_project_service.apigee]
}

resource "google_service_account" "service_account" {
  account_id   = "my-account"
  display_name = "Service Account"
}

resource "google_project_iam_binding" "synchronizer-iam" {
  project = google_project.project.project_id
  role    = "roles/apigee.synchronizerManager"
  members = [
    "serviceAccount:${google_service_account.service_account.email}",
  ]
}

resource "google_apigee_sync_authorization" "apigee_sync_authorization" {
  name       = google_apigee_organization.apigee_org.name
  identities = [
    "serviceAccount:${google_service_account.service_account.email}",
  ]
  depends_on = [google_project_iam_binding.synchronizer-iam]
}
```
