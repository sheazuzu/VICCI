# VICCI
Summary:
Create a new GO Application which provides two Ednpoints. One to recieve a list of carlines for a VICCI tenant and one to recieve a vehicle catalog overview for a given VICCI tenant and carline code
 
Requirements:
The application has to be created in GO
The application should provide the following API: vehicle_catalog_summary_service_swagger.yaml
The application should use VICCI:VIS as Datasource
The VICCI:VIS base url as well as the basic auth user and password should be configurable
 
Additional Information:
VICCI:VIS API Definition: vicci_vis_swagger.yml
VICCI:VIS Service Base URL: https://ingress.venividivicci.de/vis/
Vicci Tenants Examples:
okapi-seat-gb-en
okapi-vw-es-es
okapi-audi-de-de
Online Swagger Editor: https://editor.swagger.io/
To fetch the carlines use the VICCI:VIS endpoint /catalog/carlines with the given tenant
To create the catalog summary use the following VICCI:VIS Endpoints:
For the given tenant and carline fetch all salesgroups via /catalogue/salesgroups with all falgs set to "false"
For each salesgroup code from the previous response fetch all models via /catalogue/models using the tenant and carline code with all flags set to false
 
Solution Proposal:
 
Initiate standard GO APP (API/service/no repository)
set username/password/baseURL in .env file to communicate via api with VICCI:vis
create function to encode username/password for basic auth(base64)
create both endpoints with GET http methods
   //tenant/{vicci_tenant}/carlines and //tenant/{vicci_tenant}/catalog
to invoke second endpoint set "carline" as query parameter and "vicci_tenant" as path variable
carline and vicci_tentant will be sent to vicci api to receive salesgroups, authorization via header (base64 encoded login)
invoke /catalogue/models with given salesgroup, tenant and carline as query parameter
return recieved data and format it like defined in vicci_vis_swagger.yml
