USE [tarea4]
GO

/****** Object:  Table [dbo].[datainfo]    Script Date: 7/04/2024 23:07:19 ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[datainfo](
	[Id] [int] IDENTITY(1,1) NOT NULL,
	[Album] [nvarchar](255) NULL,
	[Year] [nvarchar](255) NULL,
	[Artist] [nvarchar](255) NULL,
	[Ranked] [nvarchar](255) NULL,
PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO


