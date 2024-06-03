using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace DbCreatorAndMigrator.Migrations
{
    /// <inheritdoc />
    public partial class PairAddLastTimeUpdateNullable : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AddColumn<DateTimeOffset>(
                name: "LastTimeUpdate",
                schema: "finance",
                table: "Pair",
                type: "timestamp with time zone",
                nullable: true);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropColumn(
                name: "LastTimeUpdate",
                schema: "finance",
                table: "Pair");
        }
    }
}
