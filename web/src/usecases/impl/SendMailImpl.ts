import SendMail from "../SendMail";
import RepositoryFactory from "../../repositories/RepositoryFactory";

class SendMailImple implements SendMail {
  private sendMailRepositry;

  constructor() {
    this.sendMailRepositry = RepositoryFactory.createSendMailRepository();
  }

  async send(args: { toAddress: string; fromAddress: string; subject: string; body: string; }): Promise<any> {
    return await this.sendMailRepositry.send(args);
  }
}

export default SendMailImple;