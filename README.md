# Spotigo

A comprehensive Go client library for the Spotify Web API that provides easy-to-use interfaces for interacting with Spotify's services.

## Features

- **Complete API Coverage**: Support for all major Spotify Web API endpoints including:
  - User Management
  - Playlist Operations
  - Track Information
  - Album Details
  - Artist Data
  - Player Controls
  - Shows and Episodes
  - Categories and Genres
  - Market Information

- **Clean Architecture**: Well-organized codebase with separate packages for different Spotify resources
- **Type Safety**: Strongly typed requests and responses
- **Error Handling**: Comprehensive error handling with meaningful error messages
- **Authentication**: Built-in support for Spotify authentication

## Installation

```bash
go get github.com/batuhannoz/spotigo
```

## Requirements

- Go 1.17 or higher
- Spotify API credentials

## Usage

Here's a quick example of how to use Spotigo:

```go
import "github.com/batuhannoz/spotigo/spotify"

// Initialize client with your Spotify token
client := spotify.New(token)

// Make API calls
// Example: Get user profile
userProfile, err := client.GetUserProfile()
if err != nil {
    // Handle error
}
```

## Package Structure

- `/album` - Album-related operations
- `/artist` - Artist information and management
- `/category` - Category listings and operations
- `/episode` - Podcast episode functionality
- `/genre` - Music genre operations
- `/market` - Market availability information
- `/player` - Playback control and state
- `/playlist` - Playlist management
- `/show` - Podcast show functionality
- `/spotify` - Core client functionality
- `/track` - Track-related operations
- `/user` - User profile and preferences

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Thanks to Spotify for providing their Web API
- All contributors who have helped with the project

## Support

If you encounter any issues or have questions, please file an issue on the GitHub repository.
