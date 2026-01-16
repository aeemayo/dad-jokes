# Dad Joke Agent 🎭

Hi there! Welcome to the Dad Joke Agent. This is a simple, fun little agent designed to brighten your day with some classic (and admittedly groan-worthy) dad jokes. 

It uses the Teneo Agent SDK and connects to OpenRouter to fetch the best (or worst?) jokes on demand.

## How to Get Started

Setting this up is super easy.

1.  **Get your keys**: You'll need a Teneo agent private key and an [OpenRouter API key](https://openrouter.ai/).
2.  **Configure environment**: Make sure you have a `.env` file in this folder. It should look something like this:

    ```env
    PRIVATE_KEY=your_teneo_private_key
    NFT_TOKEN_ID=your_nft_token_id
    OWNER_ADDRESS=your_owner_address
    OPENROUTER_API_KEY=your_openrouter_api_key_here
    ```

    > **Note**: Don't forget the `OPENROUTER_API_KEY`! The agent basically needs it to think of jokes. Without it, you'll just get a sad error message.

3.  **Run it**: Just open your terminal and run:
    ```bash
    go run .
    ```

## Usage

Once the agent is running, you can interact with it.

-   **Command**: `humor_me`
-   **What happens**: The agent reaches out to the cloud, finds a dad joke, and delivers it straight to you.

That's it! Have fun and try not to roll your eyes too hard. 😉
