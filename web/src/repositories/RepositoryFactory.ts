import SendMailRepository from "./SendMailRepository";
import SendMailRepositoryImpl from "./impl/SendMailRepositoryImpl";

class RepositoryFactory {
  private static API_ENDPOINT = "/api/v1";

  static createSendMailRepository():SendMailRepository {
    return new SendMailRepositoryImpl(this.API_ENDPOINT);
  } 
}

export default RepositoryFactory;