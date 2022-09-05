import SendMailRepository from "../SendMailRepository";
import BaseRepository from "../BaseRepository";

class SendMailRepositoryImpl extends BaseRepository implements SendMailRepository {
    send = async (args: { toAddress: string, fromAddress: string, subject: string, body: string }) => {
        const { toAddress, fromAddress, subject, body} = args;
        return this.post({ path: "/mail/send", body: {"to_address": toAddress, "from_address": fromAddress, subject, body} });
    }
}

export default SendMailRepositoryImpl;