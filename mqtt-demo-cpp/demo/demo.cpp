/*
###The MIT License

###Copyright 2020 Bosch Rexroth AG

###Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

###The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

###THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

#include <iostream>
#include <fstream>
#include <cstdlib>
#include <string>
#include <chrono>
#include <cstring>
#include <thread>
#include "mqtt/async_client.h"

//Configure variables for running the code as an application the RCU. These comments shall be removed before packaging the application in snap
const std::string DATA_FOLDER = std::getenv("SNAP_DATA");
const std::string DFLT_SERVER_ADDRESS{"ssl://127.0.0.1:8883"};

//configure variables for running the code as an application on your computer. The next 2 lines of code shall be commented before packaging the application in snap
//const std::string DATA_FOLDER{"."};
//const std::string DFLT_SERVER_ADDRESS{"ssl://192.168.188.20:8883"}; //IP Address of the RCU

const std::string DFLT_CLIENT_ID{"app_demo"};

const std::string TRUST_STORE{"/certs/boschrexroth-pjiot.pem"};
const std::string PRIVATE_KEY{"/certs/can.key"};
const std::string KEY_STORE{"/certs/can.crt"};

//topic for publishing a message that shall be sent through can2
const std::string PUB_TOPIC{"hardware/can/channel/can2/tx"};
//message be sent through can2
const std::string PUB_PAYLOAD{"(1586260384.286194) can2 18FE563D#AABBCCDDEEFF1122"};

//topic for receiving messages that are coming through can1
const std::string SUB_TOPIC_CAN_DATA{"hardware/can/channel/can1/rx"};
//topic for receiving can status messages
const std::string SUB_TOPIC_CAN_STATUS{"hardware/can/status"};

const int QOS = 1;
const auto TIMEOUT = std::chrono::seconds(30);

class callback : public virtual mqtt::callback, public virtual mqtt::iaction_listener
{
public:
	void connection_lost(const std::string &cause) override
	{
		std::cout << "\nConnection lost" << std::endl;
		if (!cause.empty())
			std::cout << "\tcause: " << cause << std::endl;
	}

	void delivery_complete(mqtt::delivery_token_ptr tok) override
	{
		std::cout << "\tDelivery complete for token: "
				  << (tok ? tok->get_message_id() : -1) << std::endl;
	}

	void message_arrived(mqtt::const_message_ptr msg) override
	{
		std::cout << "Message arrived" << std::endl;
		std::cout << "\ttopic: '" << msg->get_topic() << "'" << std::endl;
		std::cout << "\tpayload: '" << msg->to_string() << "'\n"
				  << std::endl;
	}

	void on_failure(const mqtt::token &tok) override
	{
		std::cout << "Subscription failure";
		if (tok.get_message_id() != 0)
			std::cout << " for token: [" << tok.get_message_id() << "]" << std::endl;
		std::cout << std::endl;
	}

	void on_success(const mqtt::token &tok) override
	{
		std::cout << "Subscription success";
		if (tok.get_message_id() != 0)
			std::cout << " for token: [" << tok.get_message_id() << "]" << std::endl;
		auto top = tok.get_topics();
		if (top && !top->empty())
			std::cout << "\ttoken topic: '" << (*top)[0] << "', ..." << std::endl;
		std::cout << std::endl;
	}
};

using namespace std;

int main(int argc, char *argv[])
{
	cout << "Initializing for server '" << DFLT_SERVER_ADDRESS << "'..." << endl;
	mqtt::async_client client(DFLT_SERVER_ADDRESS, DFLT_CLIENT_ID);

	callback cb;
	client.set_callback(cb);

	mqtt::ssl_options sslopts;
	sslopts.set_trust_store(DATA_FOLDER + TRUST_STORE);
	sslopts.set_key_store(DATA_FOLDER + KEY_STORE);
	sslopts.set_private_key(DATA_FOLDER + PRIVATE_KEY);

	mqtt::connect_options connopts;

	connopts.set_automatic_reconnect(true);
	connopts.set_ssl(sslopts);

	cout << "  ...OK" << endl;

	try
	{
		// Connect using SSL/TLS

		cout << "\nConnecting..." << endl;
		//mqtt::token_ptr conntok = client.connect(connopts);
		client.connect(connopts)->wait();
		cout << "Waiting for the connection..." << endl;
		cout << "  ...OK" << endl;
		client.start_consuming();
		client.subscribe(SUB_TOPIC_CAN_DATA, QOS)->wait();
		client.subscribe(SUB_TOPIC_CAN_STATUS, QOS)->wait();

		//Publish a message on the broker every 3 seconds
		while (true)
		{
			std::chrono::seconds dura( 5);
    		std::this_thread::sleep_for( dura );
			cout << "\nSending message..." << endl;
			auto msg = mqtt::make_message(PUB_TOPIC, PUB_PAYLOAD, QOS, false);
			client.publish(msg)->wait_for(TIMEOUT);
			cout << "  ...OK" << endl;
		}
	}
	catch (const mqtt::exception &exc)
	{
		cerr << exc.what() << endl;
		return 1;
	}

	return 0;
}