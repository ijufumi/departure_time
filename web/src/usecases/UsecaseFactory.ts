import SendMail from "./SendMail";
import SendMailImpl from "./impl/SendMailImpl";


class UsecaseFactory {
  static createSendMail(): SendMail {
    return new SendMailImpl();
  }
}

export default UsecaseFactory;