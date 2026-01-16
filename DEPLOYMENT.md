# Deploying Dad Joke Agent on Render

This guide will walk you through deploying your Dad Joke Agent to [Render](https://render.com/), a cloud platform that makes it easy to host apps.

## Prerequisites

1.  **Push your code to GitHub**: Make sure this code is in a GitHub repository. Render needs to see your code to build it.
2.  **Render Account**: Create an account at [render.com](https://render.com/) if you haven't already.

## Step-by-Step Deployment

### 1. Create a New Web Service or Background Worker

Since this agent runs continuously (performing tasks), you can deploy it as a **Background Worker**. However, if you plan to add a web server interface later, a **Web Service** works too.

1.  Go to your Render Dashboard.
2.  Click **New +** and select **Background Worker**.
3.  Connect your GitHub account and select your `dad-jokes` repository.

### 2. Configure the Service

Fill in the details:

-   **Name**: `dad-joke-agent` (or whatever you like)
-   **Region**: Pick the one closest to you.
-   **Runtime**: **Docker** (This is important! We created a `Dockerfile` for this).
-   **Instance Type**: The free tier (Starter) is usually fine for testing, but Background Workers might require a paid plan depending on current Render pricing.
    -   *Tip*: If you want to use the **Free Tier**, choose **Web Service** instead of Background Worker. Render's Free Web Services spin down after inactivity, but they are free.
    
### 3. Environment Variables

This is the most critical step. You need to tell Render your secrets.

Scroll down to **Environment Variables** and click **Add Environment Variable** for each of these:

| Key | Value |
| --- | --- |
| `PRIVATE_KEY` | *Your Teneo Private Key* |
| `NFT_TOKEN_ID` | *Your NFT Token ID* |
| `OWNER_ADDRESS` | *Your Owner Address* |
| `OPENROUTER_API_KEY` | *Your OpenRouter API Key* |

> **Important**: Do NOT upload your `.env` file to GitHub. Always set these secrets directly in Render.

### 4. Deploy

1.  Click **Create Background Worker** (or **Create Web Service**).
2.  Render will start building your Docker image. You'll see the logs in the dashboard.
3.  Once the build finishes, you should see "service is live" or similar.

## Troubleshooting

-   **Build Fails**: Check the "Logs" tab. If it says something about missing files, make sure you pushed everything to GitHub.
-   **App Crashes**: Check the logs. It will usually tell you if a variable like `OPENROUTER_API_KEY` is missing.

That's it! Your dad joke agent is now living in the cloud. ☁️
