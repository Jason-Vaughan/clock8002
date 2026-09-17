# Millumin Setup

This guide explains how to send a countdown from Millumin to a piClock counter.

## Configure the piClock counter

1. Open the piClock web UI.
2. Choose the timer you want the Millumin countdown to appear in.
3. Under **Time sources** > **Source** you want the countdown to appear in, enter **9** in the timer number field. Use the following timer settings:

![piClock timer settings](timerSettings.png)

## Configure Millumin

1. In Millumin, click **Interactions** (piano keys icon) > **Manage Devices** > **OSC** tab.
2. Add or select an OSC device and use the following settings. Substitute the IP address with your piClock's IP address:

![Millumin OSC settings](MilluminSettings.png)

## Troubleshooting

If the countdown is not reaching the piClock, try toggling the **Send feedback** checkbox in the Millumin OSC device settings. In at least one case this caused the OSC data to start flowing.
