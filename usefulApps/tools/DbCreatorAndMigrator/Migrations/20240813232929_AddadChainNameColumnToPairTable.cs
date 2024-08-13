using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace DbCreatorAndMigrator.Migrations
{
    /// <inheritdoc />
    public partial class AddadChainNameColumnToPairTable : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AddColumn<string>(
                name: "ChainName",
                schema: "finance",
                table: "Pair",
                type: "text",
                nullable: true);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropColumn(
                name: "ChainName",
                schema: "finance",
                table: "Pair");
        }
    }
}
