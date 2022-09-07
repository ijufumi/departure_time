import SendMail from "../SendMail";
import RepositoryFactory from "../../repositories/RepositoryFactory";

class SendMailImpl implements SendMail {
  private sendMailRepository;

  constructor() {
    this.sendMailRepository = RepositoryFactory.createSendMailRepository();
  }

  async send(args: { toAddress: string; subject: string; body: string; }): Promise<any> {
    return await this.sendMailRepository.send(args);
  }
}

export default SendMailImpl;