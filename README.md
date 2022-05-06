# Pool Toy

Pooltoy is an emoji blockchain built based on [cosmos-sdk](https://github.com/cosmos/cosmos-sdk). Nodes on pooltoy chain can mint and trading emojis.

Pooltoy can not only be run independently as a chain, but also work together  with [Slackbot](https://github.com/interchainberlin/slackbot) to send emojis on slack.

![pooltoy blockchain](./notes/cover_resize.jpg)

## Install the binary

```shell
git clone git@github.com:interchainberlin/pooltoy.git 

cd pooltoy

make install

./scripts/init.sh
```

## Start pooltoy

```shell
pooltoy start
```

Now you are ready to explore the emoji blockchain!

Open a new terminal window to try the following commands!
  
## Emoji trading

### Create new users

The account.json file contains a list of user names, addresses, and their initial emoji balances. Those data are the genesis accounts data. The user account info. can be queried as shown in the **query account info** section.

New users join in pooltoy through `pooltoy tx pooltoy create-user`. The first-created-user must be an admin user. Any account on the chain can create this first admin user because there was no admin before. From the second user creation on, only the admin can create users.
Therefore, if the first-created-user is not admin, the creating user will break. No more users can be created afterwards.

```shell
# add the key to the keyring
 pooltoy keys add [name]
 # check the key
 pooltoy keys show [name_or_address]
 # get the address
 pooltoy keys show [name_or_address] -a
 # create an admin for this key
 pooltoy tx pooltoy create-user [address] true [name] [email] --from alice -y -b block
 # create a user for this key
 pooltoy tx pooltoy create-user [address] false [name] [email] --from alice -y -b block
 #check auth info
 pooltoy q auth account [address]
 #list all the keys
 pooltoy keys list
 # list users
 pooltoy q pooltoy list-users
```

Please note:

- the above `false` in `create-user` command is for creating non-admin, true for creating an admin.
- a new created user has no emoji balances.

Presently pooltoy is designed to work together with slackbot to trade emoji in slack chat, slack controls the user authorization. So when you run pooltoy alone, you have permissions to use all the accounts on the pooltoy chain. You can play the role of admin or any other user's role. For example, you can send from any account to another if both accounts exist, and the sender has sufficient funds.

### Queries

#### query account info

```shell
# account info
pooltoy keys show [name_or_address]
# account address
pooltoy keys show [name]  -a
```

#### query balances

```shell
pooltoy query bank balances [address] -o json
```

### Transactions

```shell
pooltoy tx bank send [sender_name] [recipient_address] [amount][emoji] --from [sender_name] -y -b block
# or 
pooltoy tx bank send [from_key_or_address] [to_address] [amount][emoji] -y -b block
# -b block returns the instant transaction result
````

### Mint emoji

Pooltoy allows users to mint one emoji per 24h without any cost. Users can send this minted emoji to themselves or to other users.

```shell
# mint
pooltoy tx faucet mintfor [recipient_address] [emoji] --from [sender_name] -y -b block
# check timeleft(s) for next mint
pooltoy q faucet when-brrr [address]
```

## stop pooltoy

```shell
killall -9 pooltoy
```

## Helps

Please use `pooltoy [command] --help` to explore more commands.
