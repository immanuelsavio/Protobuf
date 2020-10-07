  
service RepositoryService {
	rpc add (Repository) returns (AddRepositoryResponse);
}
 
message AddRepositoryResponse {
	Repository addedRepository = 1;
	Error error = 2;
}
message Error {
	string code = 1;
	string message = 2;
}