import SendMailRepository from "../SendMailRepository";
import BaseRepository from "../BaseRepository";

class SendMailRepositoryImpl extends BaseRepository implements SendMailRepository {
    send = async (args: { toAddress: string, fromAddress: string, subject: string, body: string }) => {
        return this.post({ path: "/mail/send", body: args });
    }
}

export default SendMailRepositoryImpl;