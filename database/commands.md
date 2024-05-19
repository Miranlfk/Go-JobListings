# GraphQL Operations for JobListings

## Create Job Listing

The user can create a Job listing via the following command:
```
mutation CreateJobListing($input: CreateJobListingInput!){ 
  createJobListing(input:$input){ 
    _id 
    title 
    description 
    company 
    url 
  } 
}
```
Provide the relevant variables as mentioned below to successfully create a Job Listing:

`{ "input": { "title": "<job_title>", "description": "<job_description>", "company": "<company_name>", "url": "<company_url>" } }`


## Get All Job Listings

The user can retrieve all the Job Listings via the following query:
```
query GetAllJobs{ 
  jobs{ 
    _id 
    title 
    description 
    company 
    url 
  } 
}
```

## Get a Job Listing

The user can retrieve a single Job Listing by providing a job id via the following query:
```
query GetJob($id: ID!){ 
  job(id:$id){ 
    _id 
    title 
    description 
    url 
    company 
  } 
}
```
Provide the id variable as mentioned below to successfully retrieve a Job Listing:

`{"id": "<id>"}`

## Update a Job Listing

The user can update a Job listing by providing a job id via the following command:
```
mutation UpdateJob($id: ID!, $input: UpdateJobListingInput!) { 
  updateJobListing(id:$id,input:$input){ 
    title 
    description 
    _id 
    company 
    url 
  } 
}
```
Provide the id and desired input as mentioned below to successfully update a Job Listing:

` "id": "<id>", "input": { "title": "<new_job_title>", "description": "<new_job_description>", "url": "<new_company_url>" } }`

## Delete a Job Listing

The user can delete a single Job Listing by providing a job id via the following query:
```
mutation DeleteJobListing($id: ID!) { 
  deleteJobListing(id:$id){ 
    deleteJobId 
  } 
}
```

Provide the id as mentioned below to successfully delete a Job Listing:

`{"id": "<id>"}`