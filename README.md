# Azure Event Hub - Golang Producer Client

This project demonstrates how to send events to **Azure Event Hub** using the Azure SDK for Go. The authentication is done using **Managed Identity** or default credentials via the `azidentity` package.

## Prerequisites

Before running this code, you'll need:

1. **Azure Event Hub**: A namespace and Event Hub created in Azure.
2. **Microsoft Entra (Azure Active Directory)**: An **App Registration** or **Managed Identity** enabled in Azure with proper permissions to send events to Event Hub.
3. **Azure Account**: The account used for authentication must have appropriate permissions to manage and send data to the Event Hub.
4. **Install Dependencies**
5. **Running the Code**
6. **References**

## Setup

### 1. Azure Event Hub Configuration

- Create a namespace in Azure Event Hubs.
- Inside the namespace, create an Event Hub.
- Add a **Managed Identity** or register an application in **Microsoft Entra (Azure Active Directory)** with the required permissions for **Azure Event Hub** (see the reference links below).

### 2. Microsoft Entra Permissions Configuration

- Register an application in **Microsoft Entra** or use **Managed Identity** for authentication.
- In the Azure portal, go to **Microsoft Entra ID** (formerly Azure AD) and register an application. Grant the necessary permissions for the Event Hub.
- Assign the **Event Hubs Data Sender** role to the application or the Managed Identity to allow it to send events to the Event Hub.

**Detailed Documentation:**

- [How to create a service principal in Azure](https://learn.microsoft.com/en-us/entra/identity-platform/howto-create-service-principal-portal)

### 3. Environment Variables Setup

Use a `.env` file to configure the necessary environment variables if required.

Example `.env`:

```bash
# Example environment variables
AZURE_TENANT_ID=<your-tenant-id>
AZURE_CLIENT_ID=<your-client-id>
AZURE_CLIENT_SECRET=<your-client-secret>  # Not needed if using Managed Identity
```

### 4. Install Dependencies

Make sure the necessary Azure libraries are installed. If you're using Go Modules, you can install them by running:

```bash
go mod tidy
```

### 5. Running the Code

Once the project is set up, you can run the code with the following command:

```bash
go run main.go
```

### 6. References

- [Azure SDK for Go Samples](https://github.com/azure-samples/azure-sdk-for-go-samples)
- [Send events to Azure Event Hub using Go](https://learn.microsoft.com/en-us/azure/event-hubs/event-hubs-go-get-started-send)
- [How to create a service principal in Microsoft Entra](https://learn.microsoft.com/en-us/entra/identity-platform/howto-create-service-principal-portal)
