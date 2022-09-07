import SendMailRepository from "../SendMailRepository";
import BaseRepository from "../BaseRepository";

class SendMailRepositoryImpl extends BaseRepository implements SendMailRepository {
    send = async (args: { toAddress: string, subject: string, body: string }) => {
        const { toAddress, subject, body} = args;
        return this.post({ path: "/mail/send", body: {"to_address": toAddress, subject, body} });
    }
}

export default SendMailRepositoryImpl;