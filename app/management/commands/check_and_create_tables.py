from django.core.management.base import BaseCommand
from django.db import connection
from django.apps import apps
from django.core.management import call_command


class Command(BaseCommand):
    help = 'Check if model tables exist in the database and create them if not.'

    def handle(self, *args, **options):
        # Get all models from your app
        app_models = apps.get_models()
        with connection.cursor() as cursor:
            for model in app_models:
                # Construct the SQL statement to check for table existence
                table_name = model._meta.db_table
                cursor.execute("SELECT to_regclass(%s)", [table_name])
                if cursor.fetchone()[0] is None:
                    # Table doesn't exist, so we need to apply migrations
                    self.stdout.write(self.style.WARNING(
                        f'Table for model {model._meta.object_name} does not exist. Applying migrations...'))
                    call_command('migrate', app_label=model._meta.app_label, verbosity=0)
                else:
                    # Table exists, ensure the migration for this model is marked as applied
                    self.stdout.write(self.style.SUCCESS(
                        f'Table for model {model._meta.object_name} exists. Ensuring migrations are marked as applied...'))
                    call_command('migrate', app_label=model._meta.app_label, fake=True, verbosity=0)
