#!/bin/bash

echo "$VALUE" | gh secret set FORTIFY_DB_HOST -a actions -b -

echo "$VALUE" | gh secret set FORTIFY_DB_PASSWORD -a actions -b -
